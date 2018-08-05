import { AuthApi as AuthSwaggerApi } from '@/generated/swagger'

const endpoint = process.env.VUE_APP_GATEWAY_ENDPOINT

const authAPI = new AuthSwaggerApi(undefined, endpoint)

export const AuthAPI = authAPI







