import Vue from 'vue';
import VueRouter from 'vue-router';

Vue.use(VueRouter);

const ProblemCreate = () => import('@/views/problem/Create.vue');
const ProblemEdit = () => import('@/views/problem/Edit.vue');
const Problem = () => import('@/views/problem/Problem.vue');
const Contest = () => import('@/views/contest/Contest.vue');
const ContestCreate = () => import('@/views/contest/Create.vue');
const ContestEdit = () => import('@/views/contest/Edit.vue');
const ContestProblemList = () => import('@/views/contest/ProblemList.vue');
const User = () => import('@/views/general/User.vue');
const JudgeServer = () => import('@/views/general/JudgeServer.vue');
const System = () => import('@/views/general/System.vue');
const Announcement = () => import('@/views/general/Announcement.vue');
const ProblemTag = () => import('@/views/problem/Tag.vue');


const DashBoard = () => import('@/views/DashBoard.vue');
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
  { path: '/contest/problem', component: ContestProblemList },
  { path: '/general/user', component: User },
  { path: '/general/judgeServer', component: JudgeServer },
  { path: '/general/system', component: System },
  { path: '/general/announcement', component: Announcement },
  { path: '/dashboard', component: DashBoard },
  { path: '/', redirect:"/dashboard" }
  // { path: '/', component:ContestProblemList }
];

const router = new VueRouter({
  routes
});

export default router;
