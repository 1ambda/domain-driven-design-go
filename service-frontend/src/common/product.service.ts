import { ProductApi as ProductSwaggerApi } from '@/generated/swagger'

const endpoint = process.env.VUE_APP_GATEWAY_ENDPOINT

const productAPI = new ProductSwaggerApi(undefined, endpoint)

export const ProductAPI = productAPI







