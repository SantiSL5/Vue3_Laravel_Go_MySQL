<?php

namespace App\Services;

use App\Interfaces\DishTypeRepositoryInterface;
use App\Http\Requests\DishType\CreateDishType;
use App\Http\Requests\DishType\UpdateDishType;
use App\Http\Resources\DishTypeResource;

class DishTypeService 
{
    private $dishTypeRepository;

    public function __construct(DishTypeRepositoryInterface $dishTypeRepository)
    {
        $this->dishTypeRepository = $dishTypeRepository;
    }

    public function getAllDishTypesService()
    {
        return DishTypeResource::collection($this->dishTypeRepository->getAllDishTypesRepo());
    }

    public function getDishTypeByIdService($id) 
    {
        try {
            return DishTypeResource::make($this->dishTypeRepository->getDishTypeByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "DishType doesn't exist"
            ]);
        }
    }

    public function createDishTypeService(CreateDishType $request)
    {
        try {
            return DishTypeResource::make($this->dishTypeRepository->createDishTypeRepo($request));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "DishType already exist"
            ]);
        }
    }

    public function updateDishTypeService(UpdateDishType $request, $id)
    {
        try {
            DishTypeResource::make($this->dishTypeRepository->getDishTypeByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "DishType doesn't exist"
            ]);
        }

        try {
            $update = $this->dishTypeRepository->updateDishTypeRepo($request, $id);
            if ($update == 1) {
                return DishTypeResource::make($this->dishTypeRepository->getDishTypeByIdRepo($id));
            } else {
                return response()->json([
                    "Message" => "Not found"
                ], 404);
            };
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "DishType already exists"
            ]);
        }
    }

    public function deleteDishTypeService($id)
    {
        $delete = $this->dishTypeRepository->deleteDishTypeRepo($id);

        if ($delete == 1) {
            return response()->json([
                "Message" => "Deleted correctly"
            ], 200);
        } else {
            return response()->json([
                "Status" => "Not found"
            ], 404);
        } 
    }
}