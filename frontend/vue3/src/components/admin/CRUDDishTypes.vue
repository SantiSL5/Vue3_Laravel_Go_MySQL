<template>
    <h2 class="mt-2">Dish Types</h2>
    <form>
        <div class="mb-3">
            <label for="name" class="form-label">Name</label>
            <input type="text" class="form-control" id="name" v-model="state.dishtype.name" required>
        </div>
        <div class="mb-3">
            <label for="photo" class="form-label">Photo</label>
            <input type="text" class="form-control" id="photo" v-model="state.dishtype.photo" required>
        </div>
        <div class="mb-3">
            <label for="order" class="form-label">Order</label>
            <input type="number" class="form-control" id="order" v-model="state.dishtype.order" required>
        </div>
        <button type="button" class="btn btn-success mb-3" @click="createDishType()" v-if="state.formType == 'create'"
            :disabled="state.dishtype.name == '' || state.dishtype.photo == '' || state.dishtype.order == ''">Create</button>
        <div v-if="state.formType == 'update'">
            <button type="button" class="btn btn-warning mb-3" @click="updateDishType()"
                :disabled="state.dishtype.name == '' || state.dishtype.photo == ''">Update</button>
            <button type="button" class="btn btn-success mb-3 ms-3" @click="changeForm()">Back to create</button>

        </div>
    </form>


    <table class="table">
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Photo</th>
                <th>Order</th>
                <th>Operations</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="dishtype in dishtypes" :key="dishtype.id">
                <th scope="row">{{ dishtype.id }}</th>
                <td>{{ dishtype.name }}</td>
                <td>{{ dishtype.photo }}</td>
                <td>{{ dishtype.order }}</td>
                <td>
                    <button type="button" @click.stop="deleteDishType(dishtype.id)"
                        class="btn btn-danger me-3">Delete</button>
                    <button type="button" @click.stop="changeForm(dishtype)" class="btn btn-warning">Update</button>
                </td>
            </tr>
        </tbody>
    </table>
</template>

<script>
import Constant from '../../Constant';
import { useStore } from 'vuex'
import { reactive } from "vue";

export default {
    props: {
        dishtypes: Array,
    },
    setup() {
        const store = useStore();
        const state = reactive({
            dishtype: { name: "", photo: "", order: "" },
            formType: "create",
            onUpdateDishType: null,
        });

        const changeForm = (dishtype = null) => {
            if (dishtype) {
                state.formType = "update";
                state.onUpdateDishType = dishtype.id
                state.dishtype.name = dishtype.name;
                state.dishtype.photo = dishtype.photo;
                state.dishtype.order = dishtype.order;
            } else {
                state.formType = "create";
                state.onUpdateDishType = null;
                state.dishtype.name = "";
                state.dishtype.photo = "";
                state.dishtype.order = "";
            }
        }

        const createDishType = () => {
            store.dispatch("dishTypeAdmin/" + Constant.CREATE_ONE_DISHTYPE, { dishtype: state.dishtype })
        }
        const deleteDishType = (id) => {
            store.dispatch("dishTypeAdmin/" + Constant.DELETE_ONE_DISHTYPE, { id });
        }
        const updateDishType = () => {
            store.dispatch("dishTypeAdmin/" + Constant.UPDATE_ONE_DISHTYPE, { dishtype: state.dishtype, id: state.onUpdateDishType });
        }

        return { state, changeForm, createDishType, deleteDishType, updateDishType }
    },
};

</script>