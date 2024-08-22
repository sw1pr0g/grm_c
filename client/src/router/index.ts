import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
    'history': createWebHistory(),
    'routes': [
        {
            'path': '/',
            'redirect': 'login',
        },
        {
            'path': '/login',
            'name': 'login',
            'component': () => import('../views/LoginView.vue')
        },
        {
            'path': '/signup',
            'name': 'signup',
            'component': () => import('../views/SignupView.vue')
        },
        {
            'path': '/recovery',
            'name': 'recovery',
            'component': () => import('../views/RecoveryView.vue')
        }
    ]
})

export default router