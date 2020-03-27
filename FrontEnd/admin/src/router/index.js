import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const ProblemCreate = () => import('@/components/problem/Create.vue');
const Problem = () => import('@/components/problem/Problem.vue');

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
};

const routes = [
  { path: '/problem/create', component: ProblemCreate },
  { path: '/problem', component: Problem },
  // { path: '/', component:Carton }
];

const router = new VueRouter({
  routes
});

export default router;
