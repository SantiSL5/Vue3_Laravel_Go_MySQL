import { ref } from "vue";
import DishTypeService from '../services/client/DishTypeService'

export const useDishTypeInfinite = (offset = "undefined", limit = "undefined") => {
    const dishtypes = ref([])
    DishTypeService.getAllDishTypes(offset, limit)
        .then(res => dishtypes.value = res.data)
        .catch(error => console.error(error))
    return dishtypes;
};