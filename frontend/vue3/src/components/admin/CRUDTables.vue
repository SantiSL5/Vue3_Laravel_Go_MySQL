<template>
    <h2 class="mt-2">Tables</h2>
    <form>
        <div class="mb-3">
            <label for="code" class="form-label">Code</label>
            <input type="number" class="form-control" id="code" v-model="state.table.code" required>
        </div>
        <div class="mb-3">
            <label for="category" class="form-label">Category ID</label>
            <input type="number" class="form-control" id="category" v-model="state.table.category" required>
        </div>
        <div class="mb-3">
            <label for="capacity" class="form-label">Capacity</label>
            <input type="number" class="form-control" id="capacity" v-model="state.table.capacity" required>
        </div>
        <div class="mb-3">
            <label for="enabled" class="form-check-label pe-2">Enabled?</label>
            <input type="checkbox" class="form-check-input" name="enabled" id="" v-model="state.table.enabled" required>
        </div>
        <button type="button" class="btn btn-success mb-3" @click="createTable()" v-if="state.formType == 'create'"
            :disabled="state.table.code == '' || state.table.category == '' || state.table.capacity == ''">Create</button>
        <div v-if="state.formType == 'update'">
            <button type="button" class="btn btn-warning mb-3" @click="updateTable()"
                :disabled="state.table.name == '' || state.table.photo == ''">Update</button>
            <button type="button" class="btn btn-success mb-3 ms-3" @click="changeForm()">Back to create</button>

        </div>
    </form>

    <table class="table">
        <thead>
            <tr>
                <th>ID</th>
                <th>Code</th>
                <th>Category</th>
                <th>Capacity</th>
                <th>Enabled?</th>
                <th>Operations</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="table in tables" :key="table.id">
                <th scope="row">{{ table.id }}</th>
                <td>{{ table.code }}</td>
                <td>{{ table.category.id }} ({{ table.category.name }})</td>
                <td>{{ table.capacity }}</td>
                <td>
                    <input type="checkbox" name="enabled" id="" :checked="table.enabled == 1"
                        @click="updateTable(table)">
                </td>
                <td>
                    <button type="button" @click.stop="deleteTable(table.id)"
                        class="btn btn-danger me-3">Delete</button>
                    <button type="button" @click.stop="changeForm(table)" class="btn btn-warning">Update</button>
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
        tables: Array,
    },
    setup() {
        const store = useStore();
        const state = reactive({
            table: { code: "", category: "", capacity: "", enabled: false },
            formType: "create",
            onUpdateTable: null,
        });

        const changeForm = (table = null) => {
            if (table) {
                state.formType = "update";
                state.onUpdateTable = table.id
                state.table.code = table.code;
                state.table.category = table.category.id;
                state.table.capacity = table.capacity;
                state.table.enabled = table.enabled;
            } else {
                state.formType = "create";
                state.onUpdateTable = "";
                state.table.code = "";
                state.table.category = "";
                state.table.capacity = "";
                state.table.enabled = false;
            }
        }

        const createTable = () => {
            store.dispatch("tableAdmin/" + Constant.CREATE_ONE_TABLE, { table: state.table })
        }
        const deleteTable = (id) => {
            store.dispatch("tableAdmin/" + Constant.DELETE_ONE_TABLE, { id });
        }
        const updateTable = (table = null) => {
            if (table) {
                table.category = table.category.id;
                const id = table.id;
                table.enabled = !table.enabled;

                store.dispatch("tableAdmin/" + Constant.UPDATE_ONE_TABLE, { table: table, id: id });
            } else {
                store.dispatch("tableAdmin/" + Constant.UPDATE_ONE_TABLE, { table: state.table, id: state.onUpdateTable });
            }
        }

        return { state, changeForm, createTable, deleteTable, updateTable }
    },
};

</script>