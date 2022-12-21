import Constant from "../../../Constant";
import DishService from "../../../services/client/DishService";

export const dishClient = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_DISHTYPES]: (state, payload) => {
            state.disheslist = payload;
        },
        [Constant.GET_ONE_DISHTYPE]: (state) => {
            state.allUsers;
        },
    },
    actions: {
        [Constant.GET_ALL_DISHTYPES]: (store) => {
            DishService.getAllDishes().then(data => {
                store.commit(Constant.GET_ALL_DISHTYPES, data.data);
            });
        },
        [Constant.GET_ONE_DISHTYPE]: (store) => {
            DishService.getDishById().then(data => {
                store.commit(Constant.GET_ONE_DISHTYPE, data.data);
            });
        },
    },
    getters: {
        getDish(state) {
            return state.disheslist;
        }
    }
}
