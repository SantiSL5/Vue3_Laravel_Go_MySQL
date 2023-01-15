<template>
    <h2 class="mt-2">Users</h2>
    <form>
        <div class="mb-3">
            <label for="username" class="form-label">Username</label>
            <input type="text" class="form-control" id="username" v-model="state.user.username" required>
        </div>
        <div class="mb-3">
            <label for="email" class="form-label">Email</label>
            <input type="email" class="form-control" id="email" v-model="state.user.email" required>
        </div>
        <div class="mb-3">
            <label for="password" class="form-label">Password</label>
            <input type="password" class="form-control" id="password" v-model="state.user.password" required>
        </div>
        <div class="mb-3">
            <label for="enabled" class="form-check-label pe-2">Enabled?</label>
            <input type="checkbox" class="form-check-input" name="enabled" id="enabled" v-model="state.user.enabled" required checked>
        </div>
        <button type="button" class="btn btn-success mb-3" @click="createUser()" v-if="state.formType == 'create'"
            :disabled="state.user.username == '' || state.user.email == '' || state.user.password == ''">Create</button>
        <div v-if="state.formType == 'update'">
            <button type="button" class="btn btn-warning mb-3" @click="updateUser()"
            :disabled="state.user.username == '' || state.user.email == ''">Update</button>
            <button type="button" class="btn btn-success mb-3 ms-3" @click="changeForm()">Back to create</button>
        </div>
    </form>

    <table class="table">
        <thead>
            <tr>
                <th>ID</th>
                <th>Username</th>
                <th>Email</th>
                <th>Type</th>
                <th>Image</th>
                <th>Enabled?</th>
                <th>Operations</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="user in users" :key="user.id">
                <th scope="row">{{ user.id }}</th>
                <td>{{ user.username }}</td>
                <td>{{ user.email }}</td>
                <td>{{ user.type }}</td>
                <td>{{ user.image }}</td>
                <td>
                    <input type="checkbox" name="enabled" id="" :checked="user.enabled == true"
                        @click="updateUser(user)">
                </td>
                <td>
                    <button type="button" @click.stop="deleteUser(user.id)"
                        class="btn btn-danger me-3">Delete</button>
                    <button type="button" @click.stop="changeForm(user)" class="btn btn-warning">Update</button>
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
        users: Array,
    },
    setup() {
        const store = useStore();
        const state = reactive({
            user: { username: "", email: "", password: "", type: "client", enabled: true },
            formType: "create",
            onUpdateUser: null,
        });

        const changeForm = (user = null) => {
            if (user) {
                state.formType = "update";
                state.onUpdateUser = user.id
                state.user.username = user.username;
                state.user.email = user.email;
                state.user.enabled = user.enabled;
                state.user.password = "";
            } else {
                state.formType = "create";
                state.onUpdateUser = "";
                state.user.username = "";
                state.user.password = "";
                state.user.email = "";
                state.user.enabled = true;
            }
        }

        const createUser = () => {
            store.dispatch("userAdmin/" + Constant.CREATE_ONE_USER, { user: state.user })
        }
        const deleteUser = (id) => {
            store.dispatch("userAdmin/" + Constant.DELETE_ONE_USER, { id });
        }
        const updateUser = (user = null) => {
            if (user) {
                // const id = user.id;
                user.enabled = !user.enabled;

                store.dispatch("userAdmin/" + Constant.UPDATE_ONE_USER, { user: {enabled: user.enabled}, id: user.id });
            } else {
                if (state.user.password == "") {
                    delete state.user.password
                }
                store.dispatch("userAdmin/" + Constant.UPDATE_ONE_USER, { user: state.user, id: state.onUpdateUser });
            }
        }

        return { state, changeForm, createUser, deleteUser, updateUser }
    },
};

</script>