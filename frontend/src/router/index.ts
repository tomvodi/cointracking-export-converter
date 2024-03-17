import ExportView from "../views/ExportView.vue";
import SettingsView from "../views/SettingsView.vue";
import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
    {
        path: '/export',
        name: 'export',
        component: ExportView
    },
    {
        path: '/',
        redirect: '/export'
    },
    {
        path: '/settings',
        name: 'settings',
        component: SettingsView
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

export {router}
