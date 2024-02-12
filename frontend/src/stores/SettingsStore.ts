import {defineStore} from "pinia";
import {ref} from "vue";
import {common} from "../../wailsjs/go/models";


export const useSettingsStore = defineStore('settings', () => {
    const timezone = ref<string>('')
    const allTimezones = ref<Array<common.TimezoneData>>([])

    return {timezone, allTimezones}
})