import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

let proto
let mHost, mgHost, ctHost, tHost
let mPublicPort, mPrivatePort, mgPublicPort, mgPrivatePort, ctPrivatePort, tPublicPort

if (__ENV.API_GATEWAY_HOST && !__ENV.API_GATEWAY_PORT || !__ENV.API_GATEWAY_HOST && __ENV.API_GATEWAY_PORT) {
  fail("both API_GATEWAY_HOST and API_GATEWAY_PORT should be properly configured.")
}

export const apiGatewayMode = (__ENV.API_GATEWAY_HOST && __ENV.API_GATEWAY_PORT);

if (__ENV.API_GATEWAY_PROTOCOL) {
  if (__ENV.API_GATEWAY_PROTOCOL !== "http" && __ENV.API_GATEWAY_PROTOCOL != "https") {
    fail("only allow `http` or `https` for API_GATEWAY_PROTOCOL")
  }
  proto = __ENV.API_GATEWAY_PROTOCOL
} else {
  proto = "http"
}

if (apiGatewayMode) {
  // api-gateway mode
  mHost = mgHost = ctHost = tHost = __ENV.API_GATEWAY_HOST
  mPrivatePort = 3083
  mgPrivatePort = 3084
  ctPrivatePort = 3086
  mPublicPort = mgPublicPort = tPublicPort = 8080
} else {
  // direct microservice mode
  mHost = "model-backend"
  mgHost = "mgmt-backend"
  ctHost = "controller-model"
  tHost = "triton-server"
  mPrivatePort = 3083
  mgPrivatePort = 3084
  ctPrivatePort = 3086
  mPublicPort = 8083
  mgPublicPort = 8084
  tPublicPort = 8001
}

export const modelPublicHost = `${proto}://${mHost}:${mPublicPort}`;
export const modelPrivateHost = `${proto}://${mHost}:${mPrivatePort}`;
export const mgmtPublicHost = `${proto}://${mgHost}:${mgPublicPort}`;
export const mgmtPrivateHost = `${proto}://${mgHost}:${mgPrivatePort}`;
export const controllerPrivateHost = `${proto}://${ctHost}:${ctPrivatePort}`;
export const tritonPublicHost = `${proto}://${tHost}:${tPublicPort}`;

export const controllerGRPCPrivateHost = `${ctHost}:${ctPrivatePort}`;

export const modelResourcePermalink = `resources/${uuidv4()}/types/models`

export const serviceResourcePermalink = `resources/${uuidv4()}/types/services`
