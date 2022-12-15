<?php

namespace App\Interfaces;
use App\Http\Requests\DishType\CreateDishType;
use App\Http\Requests\DishType\UpdateDishType;

interface DishTypeRepositoryInterface
{
    function getAllDishTypesRepo();
    function getDishTypeByIdRepo($id);
    function createDishTypeRepo(CreateDishType $request);
    function updateDishTypeRepo(UpdateDishType $request,$id);
    function deleteDishTypeRepo($id);
}
