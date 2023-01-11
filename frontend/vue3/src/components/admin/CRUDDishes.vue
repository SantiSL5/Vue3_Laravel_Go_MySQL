<template>
    <h2 class="mt-2">Dishes</h2>
    <form>
        <div class="mb-3">
            <label for="name" class="form-label">Name</label>
            <input type="text" class="form-control" id="name" v-model="state.dish.dish" required>
        </div>
        <div class="mb-3">
            <label for="type" class="form-label">Type ID</label>
            <input type="number" class="form-control" id="type" v-model="state.dish.type" required>
        </div>
        <div class="mb-3">
            <label for="price" class="form-label">Price</label>
            <input type="number" class="form-control" id="price" v-model="state.dish.price" required>
        </div>
        <div class="mb-3">
            <label for="desc" class="form-label">Description</label>
            <input type="text" class="form-control" id="desc" v-model="state.dish.desc" required>
        </div>
        <button type="button" class="btn btn-success mb-3" @click="createDish()" v-if="state.formType == 'create'"
            :disabled="state.dish.dish == '' || state.dish.type == '' || state.dish.price == '' || state.dish.desc == ''">Create</button>
        <div v-if="state.formType == 'update'">
            <button type="button" class="btn btn-warning mb-3" @click="updateDish()"
                :disabled="state.dish.name == '' || state.dish.photo == ''">Update</button>
            <button type="button" class="btn btn-success mb-3 ms-3" @click="changeForm()">Back to create</button>

        </div>
    </form>

    <table class="table">
        <thead>
            <tr>
                <th>ID</th>
                <th>Name</th>
                <th>Type</th>
                <th>Price</th>
                <th>Description</th>
                <th>Operations</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="dish in dishes" :key="dish.id">
                <th scope="row">{{ dish.id }}</th>
                <td>{{ dish.dish }}</td>
                <td>{{ dish.type.id }} ({{ dish.type.name }})</td>
                <td>{{ dish.price }} €</td>
                <td>{{ dish.desc }}</td>
                <td>
                    <button type="button" @click.stop="deleteDish(dish.id)" class="btn btn-danger me-3">Delete</button>
                    <button type="button" @click.stop="changeForm(dish)" class="btn btn-warning">Update</button>
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
        dishes: Array,
    },
    setup() {
        const store = useStore();
        const state = reactive({
            dish: { dish: "", type: "", price: "", desc: "", photo: "mochi.jpg" },
            formType: "create",
            onUpdateDish: null,
        });

        const changeForm = (dish = null) => {
            if (dish) {
                state.formType = "update";
                state.onUpdateDish = dish.id
                state.dish.dish = dish.dish;
                state.dish.type = dish.type.id;
                state.dish.price = dish.price;
                state.dish.desc = dish.desc;
            } else {
                state.formType = "create";
                state.onUpdateDish = "";
                state.dish.dish = "";
                state.dish.type = "";
                state.dish.price = "";
                state.dish.desc = "";
            }
        }

        const createDish = () => {
            store.dispatch("dishAdmin/" + Constant.CREATE_ONE_DISH, { dish: state.dish })
        }
        const deleteDish = (id) => {
            store.dispatch("dishAdmin/" + Constant.DELETE_ONE_DISH, { id });
        }
        const updateDish = () => {
            store.dispatch("dishAdmin/" + Constant.UPDATE_ONE_DISH, { dish: state.dish, id: state.onUpdateDish });
        }

        return { state, changeForm, createDish, deleteDish, updateDish }
    },
};

</script>