import {defineStore} from "pinia";
import {ref} from "vue";

export const useSnackbarStore = defineStore('snackbar', () => {
    const visible = ref<boolean>(false)
    const text = ref<string>('')
    const color = ref<string>('')
    const timeout = ref<number>(5000)

    function showError(snackText: string) {
        text.value = snackText
        color.value = 'red-lighten-1'
        visible.value = true
    }

    function hideSnackbar() {
        visible.value = false
    }

    return {visible, text, color: color, timeout, showError, hideSnackbar}
})