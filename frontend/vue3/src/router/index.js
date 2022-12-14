import { createRouter, createWebHistory } from "vue-router";
// import AddTodo from '../views/AddTodo';
// import UpdateTodo from '../views/UpdateTodo';
// import Table from '../views/TableList';
// import TableList from '../views/TableList';
import Home from '../views/Home';
import Reservations from '../views/Reservations';
import Menu from '../views/Menu';
import Login from '../views/Login';
import Admin from '../views/Admin';
// import About from '../views/About';
import NotFound from '../views/NotFound';

const routes = [
  { path: "", redirect: { name: "home" } },
  { path: "/home", name: "home", component: Home },
  { path: "/reservations", name: "reservations", component: Reservations },
  { path: "/menu", name: "menu", component: Menu },
  { path: "/login", name: "login", component: Login },
  { path: "/admin", name: "admin", component: Admin },
  // { path:"/todos", name:"todoList", component: TodoList },
  // { path:"/table", name:"tables", component: CrudTable, children: [
  //   {
  //     path: '/list',
  //     component: ListTables,
  //   },
  //   {
  //     path: '/create',
  //     component: CreateTable,
  //   },
  //   {
  //     path: '/update',
  //     component: UpdateTables,
  //   },
  // ],},
  // { path: "/todos/add", name: "addTodo", component: AddTodo },
  // { path: "/todos/update/:id", name: "updateTodo", component: UpdateTodo },
  { path: "/:catchAll(.*)", component: NotFound },
];

// const routes = [
//   {
//     path: '/user/:id',
//     component: User,
// children: [
//   {
//     // UserProfile will be rendered inside User's <router-view>
//     // when /user/:id/profile is matched
//     path: 'profile',
//     component: UserProfile,
//   },
//   {
//     // UserPosts will be rendered inside User's <router-view>
//     // when /user/:id/posts is matched
//     path: 'posts',
//     component: UserPosts,
//   },
// ],
//   },
// ]


const router = createRouter({
  history: createWebHistory(),
  routes
});

export default router;