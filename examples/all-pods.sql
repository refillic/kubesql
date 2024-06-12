-- Returns all pods form all namespaces
SELECT * FROM pods

-- Returns all pods form namespace kube-system
SELECT * FROM pods p WHERE p.namespace = "kube-system"
SELECT p.name, len(p.spec.containers) FROM pods p WHERE p.namespace = "kube-sysem"

-- Returns all pods owned by Deployments named "deployment-1" (any namespace)
SELECT * FROM pods p JOIN deployments d ON p.parent = d.name WHERE d.name = "deployment-1"
SELECT d.namespace, d.name, p.name FROM pods p JOIN deployments d ON p.parent = d.name WHERE d.name = "deployment-1"