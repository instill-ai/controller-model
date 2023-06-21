import grpc from 'k6/net/grpc';
import {
    check,
    group
} from "k6";

import * as constant from "./const.js"

const clientPrivate = new grpc.Client();
clientPrivate.load(['proto/model/controller/v1alpha'], 'controller_service.proto');

export function CheckModelResource() {
    var httpModelResource = {
        "resource_permalink": constant.modelResourcePermalink,
        "model_state": "STATE_ONLINE"
    }

    clientPrivate.connect(constant.controllerGRPCPrivateHost, {
        plaintext: true
    });

    group("Controller API: Create model resource state in etcd", () => {
        var resCreateModelHTTP = clientPrivate.invoke('model.controller.v1alpha.ControllerPrivateService/UpdateResource', {
            resource: httpModelResource
        })

        check(resCreateModelHTTP, {
            "model.controller.v1alpha.ControllerPrivateService/UpdateResource response StatusOK": (r) => r.status === grpc.StatusOK,
            "model.controller.v1alpha.ControllerPrivateService/UpdateResource response modelResource name matched": (r) => r.message.resource.resourcePermalink == httpModelResource.resource_permalink,
        });
    });

    group("Controller API: Get model resource state in etcd", () => {
        var resGetModelHTTP = clientPrivate.invoke(`model.controller.v1alpha.ControllerPrivateService/GetResource`, {
            resource_permalink: httpModelResource.resource_permalink
        })

        check(resGetModelHTTP, {
            [`model.controller.v1alpha.ControllerPrivateService/GetResource ${httpModelResource.resource_permalink} response StatusOK`]: (r) => r.status === grpc.StatusOK,
            [`model.controller.v1alpha.ControllerPrivateService/GetResource ${httpModelResource.resource_permalink} response modelResource name matched`]: (r) => r.message.resource.resourcePermalink === httpModelResource.resource_permalink,
            [`model.controller.v1alpha.ControllerPrivateService/GetResource ${httpModelResource.resource_permalink} response modelResource state matched STATE_ONLINE`]: (r) => r.message.resource.modelState == "STATE_ONLINE",
        });
    });
}

export function CheckServiceResource() {
    var httpServiceResource = {
        "resource_permalink": constant.serviceResourcePermalink,
        "backend_state": "SERVING_STATUS_SERVING"
    }

    clientPrivate.connect(constant.controllerGRPCPrivateHost, {
        plaintext: true
    });

    group("Controller API: Create service resource state in etcd", () => {
        var resCreateServiceHTTP = clientPrivate.invoke('model.controller.v1alpha.ControllerPrivateService/UpdateResource', {
            resource: httpServiceResource
        })

        check(resCreateServiceHTTP, {
            "model.controller.v1alpha.ControllerPrivateService/UpdateResource response StatusOK": (r) => r.status === grpc.StatusOK,
            "model.controller.v1alpha.ControllerPrivateService/UpdateResource response service name matched": (r) => r.message.resource.resourcePermalink == httpServiceResource.resource_permalink,
        });
    });

    group("Controller API: Get service resource state in etcd", () => {
        var resGetServiceHTTP = clientPrivate.invoke(`model.controller.v1alpha.ControllerPrivateService/GetResource`, {
            resource_permalink: httpServiceResource.resource_permalink
        })

        check(resGetServiceHTTP, {
            [`model.controller.v1alpha.ControllerPrivateService/GetResource ${httpServiceResource.resource_permalink} response StatusOK`]: (r) => r.status === grpc.StatusOK,
            [`model.controller.v1alpha.ControllerPrivateService/GetResource ${httpServiceResource.resource_permalink} response service name matched`]: (r) => r.message.resource.resourcePermalink === httpServiceResource.resource_permalink,
            [`model.controller.v1alpha.ControllerPrivateService/GetResource ${httpServiceResource.resource_permalink} response service state matched STATE_ACTIVE`]: (r) => r.message.resource.backendState == "SERVING_STATUS_SERVING",
        });
    });
}
