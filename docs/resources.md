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
| AllComments           | Global    |
| AuthPlugin            | Global    |
| AuthProvider          | Global    |
| BackupPlugins         | Global    |
| Cluster               | Cluster   |
| Compliance            | Cluster   |
| ComplianceRunSchedule | Global    |
| ComplianceRuns        | Cluster   |
| Config                | Global    |
| CVE                   | Namespace |
| DebugLogs             | Global    |
| DebugMetrics          | Global    |
| Deployment            | Namespace |
| Detection             | Global    |
| Group                 | Global    |
| Image                 | Namespace |
| ImageComponent        | Namespace |
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
| ProbeUpload           | Global    |
| ProcessWhitelist      | Namespace |
| Risk                  | Namespace |
| Role                  | Global    |
| ScannerBundle         | Global    |
| ScannerDefinitions    | Global    |
| Secret                | Namespace |
| SensorUpgradeConfig   | Global    |
| ServiceAccount        | Namespace |
| ServiceIdentity       | Global    |
| User                  | Global    |
