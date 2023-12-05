import { createRouter, createWebHistory } from 'vue-router'
import IndexView from '../views/Index.vue'

const routes = [
  {path:'/', redirect: '/login'},
  {
    path: '/login',
    name: 'login',
    meta: {title: "登录"},
    component: () => import("@/views/user/LoginAndRegister.vue")
 },
  {
    path: '/:pathMatch(.*)*',
    name: 'error',
    component: component => import("@/views/404.vue"),
    meta: {
      requiresAuth: false
    }
  },
  {
    path: '/',
    name: 'home',
    meta: {title: "首页"},
    component: IndexView,
    children: [
      {
        path: 'vnc',
        name: 'vnc',
        meta: {title: "训练管理"},
        component: () => import("@/views/vnc/Index.vue"),
      },
      {
        path: 'md',
        name: 'md',
        meta: {title: "添加试题"},
        component: () => import("@/views/markdown/Index.vue"),
      },
      {
        path: 'manager',
        name: 'manager',
        meta: {title: "虚拟机管理"},
        component: () => import("@/views/vmmanger/Index.vue"),
      },
      {
        path: 'exam',
        name: 'exam',
        meta: {title: "试题管理"},
        component: () => import("@/views/exam/Index.vue"),
      },
      {
        path: 'profile',
        name: 'profile',
        meta: {title: "个人信息"},
        component: () => import("@/views/profile/Index.vue"),
      },
      {
        path: 'docker',
        name: 'docker',
        meta: {title: "容器管理"},
        component: () => import("@/views/docker/Index.vue"),
      }
    ]
  },
]

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  document.title = to.meta.title

  // 检查目标路由是否存在
  const foundRoute = router.getRoutes().find(route => route.path === to.path)
  if (!foundRoute) {
    // 跳转到404页面
    next({ name: 'error' })
  }

  // 添加路由守卫
  const user = JSON.parse(localStorage.getItem('user'))
  if (to.name != 'login' && user == null) {

    // 如果不是登录页面且没有认证信息，则跳转到登录页面
    next({ name: 'login' })
  }
  next()
})
export default router
