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
        </div>
    </div>
</template>

<script>
export default {
    props: {
        tableitem: Object,
    }
};
</script>

<style>
.padding-fix {
    padding: 0 !important;
}
</style>
