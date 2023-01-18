<template>
    <div class="row col-4 justify-content-center mb-5 mt-5">
        <!-- <img :src="require(`@/assets/${tableitem.photo}`)" alt="Es definitivamente una fotografía" class="img-fluid"> -->
        <img :src="require(`@/assets/mesa.png`)" alt="Es definitivamente una fotografía" class="img-fluid col-11">
        <div class="justify-content-around row col-8">
            <div class="col">
                <div class="row justify-content-center">
                    Table {{ tableitem.code }}
                </div>
            </div>
            <div class="col">
                <div class="row">
                    <div class="col-10 col-md-4 row padding-fix">
                        <img :src="require(`@/assets/icons/people.svg`)" alt="" width="16" height="16"
                            class="img-fluid">
                    </div>
                    <div class="col">
                        {{ tableitem.capacity }}
                    </div>
                    <!-- <span
                        v-bind:class="[tableitem.status === 'disable' ? 'badge bg-danger p-2 m-1' : 'badge bg-success p-2 m-1']"
                        @click.stop="viewCategory(tableitem.id)"> {{ tableitem.status }} </span> -->
                </div>
            </div>
            <div class="row justify-content-center mt-2 ms-5">
                <h6 class="col-7">Category: {{ tableitem.category }}</h6>
            </div>
            <div class="row justify-content-center">
                <button v-if="tableitem.enabled == true" type="button" class="btn btn-light col-6"
                    @click.stop="reserve(tableitem.id)">Reserve</button>
                <button v-if="tableitem.enabled == false" type="button" class="btn btn-success col-6"
                    @click.stop="reserve(tableitem.id)" disabled>Reserved</button>
            </div>
        </div>
    </div>
</template>

<script>
import { getCurrentInstance } from 'vue';

export default {
    props: {
        tableitem: Object,
    },
    setup() {
        const { emit } = getCurrentInstance();

        const reserve = (table) => {
            emit('reservetable', table)
        }

        return { reserve };
    },

};
</script>

<style>
.padding-fix {
    padding: 0 !important;
}
</style>
