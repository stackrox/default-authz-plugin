# Resources

The following table lists all resources currently supported by the
StackRox Kubernetes Security Platform. Resources correspond to `noun`s
in the authorization plugin API. The `Scope` column indicates
whether this resource is scoped to a cluster or even a cluster/namespace
pair, or whether it is a global resource (i.e., attributes will always be empty).

| Resource              | Scope     |
| --------------------- | --------- |
| APIToken              | Global    |
| Alert                 | Namespace |
| AuthPlugin            | Global    |
| AuthProvider          | Global    |
| BackupPlugins         | Global    |
| Cluster               | Cluster   |
| Compliance            | Cluster   |
| ComplianceRunSchedule | Global    |
| ComplianceRuns        | Cluster   |
| Config                | Global    |
| DebugLogs             | Global    |
| DebugMetrics          | Global    |
| Deployment            | Namespace |
| Detection             | Global    |
| Group                 | Global    |
| Image                 | Namespace |
| ImageIntegration      | Global    |
| ImbuedLogs            | Global    |
| Indicator             | Namespace |
| K8sRole               | Namespace |
| K8sRoleBinding        | Namespace |
| K8sSubject            | Namespace |
| Licenses              | Global    |
| Namespace             | Namespace |
| NetworkGraph          | Namespace |
| NetworkPolicy         | Namespace |
| Node                  | Cluster   |
| Notifier              | Global    |
| Policy                | Global    |
| ProcessWhitelist      | Namespace |
| Role                  | Global    |
| ScannerDefinitions    | Global    |
| Secret                | Namespace |
| ServiceAccount        | Namespace |
| ServiceIdentity       | Global    |
| User                  | Global    |
