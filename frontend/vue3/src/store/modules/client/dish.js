import Constant from "../../../Constant";
import DishService from "../../../services/client/DishService";

export const dishClient = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.GET_ALL_DISHES]: (state, payload) => {
            state.disheslist = payload;
        },
        [Constant.GET_ONE_DISH]: (state) => {
            state.allUsers;
        },
    },
    actions: {
        [Constant.GET_ALL_DISHES]: (store) => {
            DishService.getAllDishes().then(data => {                
                store.commit(Constant.GET_ALL_DISHES, data.data);
            });
        },
        [Constant.GET_ONE_DISH]: (store) => {
            DishService.getDishById().then(data => {
                store.commit(Constant.GET_ONE_DISH, data.data);
            });
        },
    },
    getters: {
        getDish(state) {
            return state.disheslist;
        }
    }
}
