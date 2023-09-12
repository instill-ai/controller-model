import { uuidv4 } from 'https://jslib.k6.io/k6-utils/1.4.0/index.js';

let proto

export const apiGatewayMode = (__ENV.API_GATEWAY_URL && true)

if (__ENV.API_GATEWAY_PROTOCOL) {
  if (__ENV.API_GATEWAY_PROTOCOL !== "http" && __ENV.API_GATEWAY_PROTOCOL != "https") {
    fail("only allow `http` or `https` for API_GATEWAY_PROTOCOL")
  }
  proto = __ENV.API_GATEWAY_PROTOCOL
} else {
  proto = "http"
}

export const controllerGRPCPrivateHost = "controller-model:3086"

export const modelResourcePermalink = `resources/${uuidv4()}/types/models`

export const serviceResourcePermalink = `resources/${uuidv4()}/types/services`
