[![Build status](https://api.travis-ci.org/nuxeo/etcd-netfw.svg)](https://travis-ci.org/nuxeo/etcd-netfw)

etcd-netfw
==========

etcd-netfw is a dynamic TCP proxy (net forwarder) which configuration is based on [etcd](https://github.com/coreos/etcd). It provides real time dynamic reconfiguration of destination hosts and is able to do some load balancing.

It is part of the nuxeo.io infrastructure.


How it works
------------

It is a TCP proxy which fetch its configuration from `etcd`. The configuration format is the following. Let's says the `servicePath` parameter is set to `/services/postgres/` and given the following `etcd` values :

    /services/postgres/1: {"host": "172.31.51.34", "port": 5432}
    /services/postgres/2: {"host": "172.31.51.35", "port": 5432}

Then, when asking for a connection to `etcd-netfw`, it will forward alternatively to the first and second host.

Configuration
-------------

Several parameters allow to configure the way the proxy behave :

 * `acceptAddr` allows to setup a listening address (0.0.0.0/1337 by default)
 * `etcdAddress` address of a list of etcd nodes (coma separated)
 * `servicePath` : path of the key where the service instances are resistered.

Use in docker
-------------

`etcd-netfw` is used in [nuxeo.io](https://github.com/nuxeo/nuxeo.io) to build dynamic Docker ambassadors. A small explanation can be found here : [https://github.com/nuxeo/nuxeo.io-scripts/tree/master/docker/service-amb](https://github.com/nuxeo/nuxeo.io-scripts/tree/master/docker/service-amb)

How to build
------------

We use [GOM](https://github.com/mattn/gom) to build.

    go get github.com/mattn/gom
    gom install


Report & Contribute
-------------------

We are glad to welcome new developers on this initiative, and even simple usage feedback is great.
- Ask your questions on [Nuxeo Answers](http://answers.nuxeo.com)
- Report issues on this github repository (see [issues link](http://github.com/nuxeo/etcd-netfw/issues) on the right)
- Contribute: Send pull requests!


About Nuxeo
-----------

Nuxeo provides a modular, extensible Java-based
[open source software platform for enterprise content management](http://www.nuxeo.com/en/products/ep),
and packaged applications for [document management](http://www.nuxeo.com/en/products/document-management),
[digital asset management](http://www.nuxeo.com/en/products/dam) and
[case management](http://www.nuxeo.com/en/products/case-management).

Designed by developers for developers, the Nuxeo platform offers a modern
architecture, a powerful plug-in model and extensive packaging
capabilities for building content applications.

More information on: <http://www.nuxeo.com/>
