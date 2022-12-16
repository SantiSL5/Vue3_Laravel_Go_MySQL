<?php

namespace App\Services;

use App\Interfaces\DishRepositoryInterface;
use App\Interfaces\DishTypeRepositoryInterface;
use App\Http\Requests\Dish\CreateDish;
use App\Http\Requests\Dish\UpdateDish;
use App\Http\Resources\DishResource;
use App\Http\Resources\DishTypeResource;

class DishService
{
    private $dishRepository;
    private $dishTypeRepository;

    public function __construct(DishRepositoryInterface $dishRepository, DishTypeRepositoryInterface $dishTypeRepository)
    {
        $this->dishRepository = $dishRepository;
        $this->dishTypeRepository = $dishTypeRepository;
    }

    public function getAllDishesService()
    {
        return DishResource::collection($this->dishRepository->getAllDishesRepo());
    }

    public function getDishByIdService($id)
    {
        try {
            return DishResource::make($this->dishRepository->getDishByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Dish doesn't exist"
            ]);
        }
    }

    public function createDishService(CreateDish $request)
    {
        try {
            DishTypeResource::make($this->dishTypeRepository->getDishTypeByIdRepo($request->get('type')));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "DishType doesn't exist"
            ]);
        }

        try {
            return DishResource::make($this->dishRepository->createDishRepo($request));
        } catch (\Exception $e) {
            return response()->json("Dish already exists");
        }
    }

    public function updateDishService(UpdateDish $request, $id)
    {
        try {
            DishResource::make($this->dishRepository->getDishByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Dish doesn't exist"
            ]);
        }

        if ($request->get('type')) {
            try {
                DishTypeResource::make($this->dishTypeRepository->getDishTypeByIdRepo($request->get('type')));
            } catch (\Exception $e) {
                return response()->json([
                    "Message" => "DishType doesn't exist"
                ]);
            }
        }

        try {
            $update = $this->dishRepository->updateDishRepo($request, $id);
            if ($update == 1) {
                return response()->json([
                    "Message" => "Updated correctly"
                ]);
            } else {
                return response()->json([
                    "Message" => "Not found"
                ], 404);
            };
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Dish already exists"
            ]);
        }
    }

    public function deleteDishService($id)
    {
        $delete = $this->dishRepository->deleteDishRepo($id);

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
