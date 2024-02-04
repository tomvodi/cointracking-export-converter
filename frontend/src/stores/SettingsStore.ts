import {defineStore} from "pinia";
import {ref} from "vue";


export const useSettingsStore = defineStore('settings', () => {
    const timezone = ref<string>('')

    return {timezone}
})