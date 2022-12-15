<?php

namespace App\Repositories;
use App\Models\DishType;
use App\Interfaces\DishTypeRepositoryInterface;
use App\Http\Requests\DishType\CreateDishType;
use App\Http\Requests\DishType\UpdateDishType;
use Illuminate\Database\Eloquent\Builder;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

class DishTypeRepository implements DishTypeRepositoryInterface
{

    public function getAllDishTypesRepo()
    {
        return DishType::all();
    }

    public function getDishTypeByIdRepo($id) 
    {
        return DishType::where('id', $id)->firstOrFail();
    }

    public function createDishTypeRepo(CreateDishType $request)
    {
        return DishType::create($request->validated());
    }

    public function updateDishTypeRepo(UpdateDishType $request, $id)
    {
        return DishType::where('id', $id)->update($request->validated());
    }

    public function deleteDishTypeRepo($id)
    {
        return DishType::where('id', $id)->delete();
    }
}
