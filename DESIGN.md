== Design ==

== TenantEgressPolicy ==
- dnsName
- cidrSelector
- groupRef
- Allow ONLY

== AdminEgressPolicy ==
- beforePolicies
  - dnsName
  - cidrSelector
- afterPolicies
  - deny all possibility

== AdminGroupEgressPolicy ==
- list of
  - dnsName
  - cidrSelector

Reconcile to EgressFirewall or EgressNetworkPolicy - depending what is present on cluster
- Both if they are both
- Don't update when there is no change
- Update when there is a change
- On start update all