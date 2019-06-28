# Overview

This document provides an overview for the Default Authorization Plugin
for the StackRox Kubernetes Security Platform. This plugin provides a
simple, rule-based configuration mechanism to enforce fine-grained access
controls in the StackRox Platform, as well as serving as a reference 
implementation for customers wishing to implement their custom authorization
plugins.

## Purpose

An Authorization Plugin allows to externalize access control logic in
the StackRox Kubernetes Security Platform: whenever a user -- human or
machine-based -- attempts to access a resource in StackRox and an
Authorization Plugin is configured, a REST call is made to a specified
HTTP(S) endpoint of the Authorization Plugin to determine whether
or not this user should be granted access to the given resource or
not. The possible granularity of access checks exceeds that which is
natively offered by the StackRox Kubernetes Security Platform: in
addition to specifying the type of access (view or modify) and the
resource being accessed, the Authorization Plugin can grant or deny an
access based on which cluster or namespace the resource is associated
with (if applicable -- a full list of resources, along with their
respective scopes can be found [here](resources.md)).

## Concepts

The interaction between the StackRox Kubernetes Security Platform and an
Authorization Plugin consists in StackRox querying the plugin whether a
*principal* is allowed access to a certain *(access) scope*. These concepts
can be defined as follows.

### Principals
A **principal** is an entity accessing a StackRox API. This can be either
a human user (identified through an external Identity Provider), or an API
token configured inside of the StackRox Kubernetes Security Platform. 

For the purposes of Authorization Plugins, these are the only relevant principals.
Specifically, for both StackRox services (such as Sensor) as well as 
administrator users authenticating via HTTP Basic authentication, the
Authorization Plugin will not be queried.

To an Authorization Plugin, a principal is identified by its method of authentication
(the *Authentication Provider*) as well as its *attributes*. See the
[API documentation](api.md) for a detailed description on how a principal
is communicated to an Authorization Plugin.

### Access Scopes
An **access scope** (or **scope**) describes access to a certain resource in
a certain way. The mode of access (`view` or `edit`) is referred to as the *verb*,
and the [resource](resources.md) to be accessed is also referred to as the *noun*.

The following are examples for access scopes, expressed in natural language:
- View all deployments
- Modify image integrations
- View all namespaces in cluster `remote`
- Modify process whitelists in namespace `test` in cluster `remote`

It is worth noting that scopes can be nested. For example, "View all namespaces
in cluster `remote`" is a sub-scope of "View all namespaces (in all clusters)". If
the Authorization Plugin grants access to a certain scope, the StackRox Kubernetes Security
Platform assumes that access to all sub-scopes of this scope are accessible as well.

The StackRox Kubernetes Security Platform will generally query for more general scopes
first, and only if access to these scopes is denied, will query for access to sub-scopes
of this scope. However, there is no fixed sequence in which this occurs: for example, if
access for viewing all deployments is denied, the next query might be for access to view
deployments in a certain namespace, without ever querying access for all deployments in the
enclosing cluster. For this reason, it is advisable that Authorization Plugins respect scope
hierarchies as well, i.e., if access to a scope is granted, access to all
sub-scopes of this scopes should be granted as well.

## Default Authorization Plugin

The Default Authorization Plugin is the authorization plugin contained in this
software bundle. It can be used as-is, or as a reference implementation for customers
who wish to write their own Authorization Plugins.

The following topics are covered in this documentation:
- [Building the plugin](building.md)
- [Configuring the HTTP server](server-config.md)
- [Description of the HTTP API](api.md)
- [Writing access rules](writing-gval-rules.md)
- [List of all resources](resources.md)
- [Deploying the plugin in Kubernetes](deploying.md)
