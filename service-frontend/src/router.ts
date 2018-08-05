import Vue from 'vue'
import VueRouter from 'vue-router'

import Login from './views/Login.vue'
import Register from './views/Register.vue'
import Home from './views/Home.vue'
import About from './views/About.vue'
import Product from './views/Product.vue'
import ProductDetail from './views/ProductDetail.vue'

Vue.use(VueRouter)

export default [
    {
        path: '/',
        name: 'home',
        displayName: 'G Street',
        component: Home,
        meta: { requiresAuth: true, common: true, left: true, },
    },
    {
        path: '/register',
        name: 'register',
        displayName: 'Register',
        component: Register,
        meta: { requiresAuth: false, common: false, left: true, },
    },
    {
        path: '/login',
        name: 'login',
        displayName: 'Login',
        component: Login,
        meta: { requiresAuth: false, common: false, left: true, },
    },
    {
        path: '/product',
        name: 'product',
        displayName: 'Product',
        component: Product,
        meta: { requiresAuth: true, common: false, left: true, },
    },
    {
        path: '/about',
        name: 'about',
        displayName: 'About',
        component: About,
        meta: { requiresAuth: false, common: false, left: false, },
    },
    {
        path: '/product/:productID',
        name: 'product.detail',
        component: ProductDetail,
        meta: { requiresAuth: true, common: false, left: false, },
    },
]

