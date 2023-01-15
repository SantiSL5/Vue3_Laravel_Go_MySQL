import { createRouter, createWebHistory } from "vue-router";
import Home from '../views/Home';
import Reservations from '../views/Reservations';
import Menu from '../views/Menu';
import Login from '../views/Login';
import Admin from '../views/Admin';
import NotFound from '../views/NotFound';

import AuthGuards from '../services/guards/AuthGuards';

const routes = [
  { path: "", redirect: { name: "home" } },
  { path: "/home", name: "home", component: Home },
  { path: "/reservations", name: "reservations", component: Reservations },
  { path: "/menu", name: "menu", component: Menu },
  { path: "/login", name: "login", component: Login, beforeEnter: AuthGuards.noAuthGuard, meta: { requiresAuth: true } },
  { path: "/admin", name: "admin", component: Admin, beforeEnter: AuthGuards.authGuardAdmin, meta: { requiresAuth: true } },
  { path: "/:catchAll(.*)", component: NotFound },
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;