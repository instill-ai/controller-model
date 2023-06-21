import grpc from 'k6/net/grpc';
import {
  check,
  group
} from 'k6';

import * as constant from "./const.js"
import * as controller_service from './controller-private.js';
const client = new grpc.Client();

client.load(['proto/model/controller/v1alpha'], 'controller_service.proto');

export let options = {
  setupTimeout: '10s',
  insecureSkipTLSVerify: true,
  thresholds: {
    checks: ["rate == 1.0"],
  },
};

export default function (data) {

  /*
   * Controller API - API CALLS
   */
  if (!constant.apiGatewayMode) {
    // Health check
    group("Controller API: Health check", () => {
      client.connect(constant.controllerGRPCPrivateHost, {
        plaintext: true
      });

      check(client.invoke('model.controller.v1alpha.ControllerPrivateService/Liveness', {}), {
        'Liveness Status is OK': (r) => r && r.status === grpc.StatusOK,
        'Response status is SERVING_STATUS_SERVING': (r) => r && r.message.healthCheckResponse.status === "SERVING_STATUS_SERVING",
      });

      check(client.invoke('model.controller.v1alpha.ControllerPrivateService/Readiness', {}), {
        'Readiness Status is OK': (r) => r && r.status === grpc.StatusOK,
        'Response status is SERVING_STATUS_SERVING': (r) => r && r.message.healthCheckResponse.status === "SERVING_STATUS_SERVING",
      });
      client.close();
    });

    controller_service.CheckModelResource()
    controller_service.CheckServiceResource()
  } else {
    console.log("No Public APIs")
  }

}

export function teardown(data) {
  if (!constant.apiGatewayMode) {
    client.connect(constant.controllerGRPCPrivateHost, {
      plaintext: true
    });
    group("Controller API: Delete all resources created by the test", () => {

      check(client.invoke(`model.controller.v1alpha.ControllerPrivateService/DeleteResource`, {
        resource_permalink: constant.modelResourcePermalink
      }), {
        [`model.controller.v1alpha.ControllerPrivateService/DeleteResource ${constant.modelResourcePermalink} response StatusOK`]: (r) => r.status === grpc.StatusOK,
      });

      check(client.invoke(`model.controller.v1alpha.ControllerPrivateService/DeleteResource`, {
        resource_permalink: constant.serviceResourcePermalink
      }), {
        [`model.controller.v1alpha.ControllerPrivateService/DeleteResource ${constant.serviceResourcePermalink} response StatusOK`]: (r) => r.status === grpc.StatusOK,
      });
    });
    client.close();
  } else {
    console.log("No Public APIs")
  }

}
