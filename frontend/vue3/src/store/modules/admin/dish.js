import Constant from "../../../Constant";
import DishService from "../../../services/admin/DishService";
import { createToaster } from "@meforma/vue-toaster"; // Debería estar en el componente

export const dishAdmin = {
    namespaced: true,

    state: {

    },
    mutations: {
        [Constant.CREATE_ONE_DISH]: (state, payload) => {
            state.disheslist.push({ ...payload });
        },
        [Constant.GET_ALL_DISHES]: (state, payload) => {
            state.disheslist = payload;
        },
        [Constant.GET_ONE_DISH]: (state) => {
            state.allUsers;
        },
        [Constant.UPDATE_ONE_DISH]: (state, payload) => {
            let index = state.disheslist.findIndex(
                item => item.id == payload.id
            );
            state.disheslist[index] = payload;
        },
        [Constant.DELETE_ONE_DISH]: (state, payload) => {
            let index = state.disheslist.findIndex(
                (item) => item.id === payload
            );
            state.disheslist.splice(index, 1);
        },
    },
    actions: {
        [Constant.CREATE_ONE_DISH]: (store, payload) => {
            DishService.createDish(payload.dish).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.statusText == "Created") {
                    toaster.success(`Dish created successfully`);
                    store.commit(Constant.CREATE_ONE_DISH, data.data);
                } else {
                    toaster.error(data.data.Message);
                }
            });
        },
        [Constant.UPDATE_ONE_DISH]: (store, payload) => {
            DishService.updateDish(payload.dish, payload.id).then(data => {
                const toaster = createToaster({
                    position: "top"
                });
                if (data.data.type) {
                    toaster.success(`Dish updated successfully`);
                    store.commit(Constant.UPDATE_ONE_DISH, data.data)
                } else {
                    toaster.error(`Dish already exists`);
                }
            });
        },
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
        [Constant.DELETE_ONE_DISH]: (store, payload) => {
            DishService.deleteDishById(payload.id).then(
                store.commit(Constant.DELETE_ONE_DISH, payload.id)
            );
        },
    },
    getters: {
        getDish(state) {
            return state.disheslist;
        }
    }
}
