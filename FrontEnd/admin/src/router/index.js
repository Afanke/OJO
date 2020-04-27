import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const ProblemCreate = () => import('@/components/problem/Create.vue');
const ProblemEdit = () => import('@/components/problem/Edit.vue');
const Problem = () => import('@/components/problem/Problem.vue');
const Contest = () => import('@/components/contest/Contest.vue');
const ContestCreate = () => import('@/components/contest/Create.vue');
const ContestEdit = () => import('@/components/contest/Edit.vue');
const Carton = () => import('@/components/problem/Carton.vue');
const User = () => import('@/components/general/User.vue');
const JudgeServer = () => import('@/components/general/JudgeServer.vue');
const System = () => import('@/components/general/System.vue');
const Announcement = () => import('@/components/general/Announcement.vue');
const ProblemTag = () => import('@/components/problem/Tag.vue');
const Test = () => import('@/components/Test.vue');
const DashBoard = () => import('@/components/DashBoard.vue');
// const Test2 = () => import('@/components/Test2.vue');

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
};

const routes = [
  { path: '/problem/create', component: ProblemCreate },
  { path: '/problem/edit', component: ProblemEdit },
  { path: '/problem/tag', component: ProblemTag },
  { path: '/problem', component: Problem },
  { path: '/contest', component: Contest },
  { path: '/contest/create', component: ContestCreate },
  { path: '/contest/edit', component: ContestEdit },
  { path: '/carton', component: Carton },
  { path: '/general/user', component: User },
  { path: '/general/judgeServer', component: JudgeServer },
  { path: '/general/system', component: System },
  { path: '/general/announcement', component: Announcement },
  { path: '/dashboard', component: DashBoard },
  // { path: '/', component:Test }
  { path: '/', component:User }
];

const router = new VueRouter({
  routes
});

export default router;
