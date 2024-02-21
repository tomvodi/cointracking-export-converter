import {defineStore} from "pinia";
import {computed, ref} from "vue";
import {common} from "../../wailsjs/go/models";


export const useSettingsStore = defineStore('settings', () => {
    const timezone = ref<string>('')
    const allTimezones = ref<Array<common.TimezoneData>>([])

    const timezoneEmpty = computed(() => {
        return timezone.value.length == 0
    })

    return {timezone, timezoneEmpty, allTimezones}
})