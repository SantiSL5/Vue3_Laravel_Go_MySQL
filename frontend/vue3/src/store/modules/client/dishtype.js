import Constant from "../../../Constant";
import DishTypeService from "../../../services/client/DishTypeService";

export const dishtypeClient = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_DISHTYPES]: (state, payload) => {
            state.dishtypeslist = payload;
        },
        [Constant.GET_ONE_DISHTYPE]: (state) => {
            state.allUsers;
        },
    },
    actions: {
        [Constant.GET_ALL_DISHTYPES]: (store) => {
            DishTypeService.getAllDishTypes().then(data => {
                store.commit(Constant.GET_ALL_DISHTYPES, data.data);
            });
        },
        [Constant.GET_ONE_DISHTYPE]: (store) => {
            DishTypeService.getDishTypeById().then(data => {
                store.commit(Constant.GET_ONE_DISHTYPE, data.data);
            });
        },
    },
    getters: {
        getDishType(state) {
            return state.dishtypeslist;
        }
    }
}
