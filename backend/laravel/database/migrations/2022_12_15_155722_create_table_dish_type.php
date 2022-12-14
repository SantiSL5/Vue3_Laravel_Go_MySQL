<?php

use Illuminate\Database\Migrations\Migration;
use Illuminate\Database\Schema\Blueprint;
use Illuminate\Support\Facades\Schema;

return new class extends Migration
{
    /**
     * Run the migrations.
     *
     * @return void
     */
    public function up()
    {
        Schema::create('dish_types', function (Blueprint $table) {
            $table->id();
            $table->timestamps();
            $table->string("name")->unique();
            $table->string("photo");
            $table->integer("order");
        });
    }
                                                                        
    /**
     * Reverse the migrations.
     *
     * @return void
     */
    public function down()
    {
        Schema::dropIfExists('dish_type');
    }
};
