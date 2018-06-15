Clearwater Sprout Cluster Operator
==================================

A demonstration of using the [Operator SDK](https://github.com/operator-framework/operator-sdk)
to create an operator to manage a cluster of Sprout (SIP router) and Bono (SIP edge proxy) nodes
for the [Clearwater IMS](http://www.projectclearwater.org/).

In this demonstration we are creating a new custom resource for Kubernetes called a
`SproutCluster`, which supports a number of shards and scale per shard. When you apply it to
Kubernetes, behind the scenes the operator will make sure to create the proper deployments for
all the Sprout pods as well as one Bono edge proxy pod per cluster. The resource also gets updated
with the pod names. Finally, if you delete the resource, all spawned resources (deployments and
pods) would also be deleted.

(Actually, Clearwater doesn't support sharding per se: we are implementing it here for the purpose
of demonstration.)
