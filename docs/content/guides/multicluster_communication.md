---
title: Multicluster traffic
description: Establish traffic policies across multiple service meshes
weight: 50
---

In the [previous guides]({{% versioned_link_path fromRoot="/guides/federate_identity/" %}}), we federated multiple meshes and established a [shared root CA for a shared identity]({{% versioned_link_path fromRoot="/guides/federate_identity/#understanding-the-shared-root-process" %}}) domain. Now that we have a logical [VirtualMesh]({{% versioned_link_path fromRoot="/reference/api/github.com.solo-io.gloo-mesh.api.networking.v1.virtual_mesh/" %}}), we need a way to establish **traffic** policies across the multiple meshes, without treating each of them individually. Gloo Mesh helps by establishing a single, unified API that understands the logical VirtualMesh construct.

## Before you begin
To illustrate these concepts, we will assume that:

* There are two clusters managed by Gloo Mesh named `cluster-1` and `cluster-2`. 
* Istio is [installed on both client clusters]({{% versioned_link_path fromRoot="/guides/installing_istio" %}})
* The `bookinfo` app is [installed across the two clusters]({{% versioned_link_path fromRoot="/guides/#bookinfo-deployed-on-two-clusters" %}})

{{% notice note %}}
Be sure to review the assumptions and satisfy the pre-requisites from the [Guides]({{% versioned_link_path fromRoot="/guides" %}}) top-level document.
{{% /notice %}}

Also ensure you have the correct context names set in your environment:

```shell
CONTEXT_1=cluster_1's_context_here
```

## Controlling cross-cluster traffic

We will now perform a *multi-cluster traffic split*, splitting traffic from the `productpage` service in `cluster-1` between `reviews-v1`, `reviews-v2`, and `reviews-v3` running in `cluster-2`.

{{< tabs >}}
{{< tab name="YAML file" codelang="yaml">}}

apiVersion: networking.mesh.gloo.solo.io/v1
kind: TrafficPolicy
metadata:
  namespace: gloo-mesh
  name: simple
spec:
  destinationSelector:
  - kubeServiceRefs:
      services:
        - clusterName: cluster-1
          name: reviews
          namespace: bookinfo
  policy:
    trafficShift:
      destinations:
        - kubeService:
            clusterName: cluster-2
            name: reviews
            namespace: bookinfo
          weight: 75
        - kubeService:
            clusterName: cluster-1
            name: reviews
            namespace: bookinfo
            subset:
              version: v1
          weight: 15
        - kubeService:
            clusterName: cluster-1
            name: reviews
            namespace: bookinfo
            subset:
              version: v2
          weight: 10
{{< /tab >}}
{{< tab name="CLI inline" codelang="shell" >}}
kubectl apply --context $CONTEXT_1 -f - << EOF
apiVersion: networking.mesh.gloo.solo.io/v1
kind: TrafficPolicy
metadata:
  namespace: gloo-mesh
  name: simple
spec:
  destinationSelector:
  - kubeServiceRefs:
      services:
        - clusterName: cluster-1
          name: reviews
          namespace: bookinfo
  policy:
    trafficShift:
      destinations:
        - kubeService:
            clusterName: cluster-2
            name: reviews
            namespace: bookinfo
          weight: 75
        - kubeService:
            clusterName: cluster-1
            name: reviews
            namespace: bookinfo
            subset:
              version: v1
          weight: 15
        - kubeService:
            clusterName: cluster-1
            name: reviews
            namespace: bookinfo
            subset:
              version: v2
          weight: 10
EOF
{{< /tab >}}
{{< /tabs >}}

Once you apply this resource to `cluster-1`, you should occasionally see traffic being routed to the reviews-v3 service, which will produce red-colored stars on the product page.

To go into slightly more detail here: The above TrafficPolicy says that:

* Any traffic destined for the *reviews service* in the *management plane cluster* should be split across several different destinations
  * 25% of traffic gets split between the v1 and v2 instances of the reviews service in the management plane cluster
  * 75% of traffic gets sent to the v3 instance of the reviews service on the remote cluster

We have successfully set up multi-cluster traffic across our application! Note that this has been done transparently to the application. The application can continue talking to what it believes is the local instance of the service. Behind the scenes we have instead routed that traffic to an entirely different cluster. 

Note that this is interesting in its own right, that we have easily achieved multi-cluster communication, but also in other scenarios like in fast disaster recovery: We can quickly route traffic to healthy instances of a service in an entirely different data center.

## See it in action

Check out "Part Four" of the ["Dive into Gloo Mesh" video series](https://www.youtube.com/watch?v=4sWikVELr5M&list=PLBOtlFtGznBjr4E9xYHH9eVyiOwnk1ciK)
(note that the video content reflects Gloo Mesh <b>v0.6.1</b>):

<iframe width="560" height="315" src="https://www.youtube.com/embed/HAr1Mw1bxB4" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>

