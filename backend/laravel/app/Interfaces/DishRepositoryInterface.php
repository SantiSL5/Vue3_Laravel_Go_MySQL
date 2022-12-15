<?php

namespace App\Interfaces;
use App\Http\Requests\Dish\CreateDish;
use App\Http\Requests\Dish\UpdateDish;

interface DishRepositoryInterface
{
    function getAllDishesRepo();
    function getDishByIdRepo($id);
    function getDishByCodeRepo($code);
    function createDishRepo(CreateDish $request);
    function updateDishRepo(UpdateDish $request,$id);
    function deleteDishRepo($id);
}
