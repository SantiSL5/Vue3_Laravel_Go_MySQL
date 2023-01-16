<?php

namespace App\Services;

use App\Interfaces\ReserveRepositoryInterface;
use App\Interfaces\TableRepositoryInterface;
use App\Interfaces\UserRepositoryInterface;
use App\Http\Requests\Reserve\CreateReserve;
use App\Http\Requests\Reserve\UpdateReserve;
use App\Http\Resources\ReserveResource;
use App\Http\Resources\TableResource;
use App\Http\Resources\UserResource;
use Carbon\Carbon;

class ReserveService
{
    private $reserveRepository;
    private $tableRepository;
    private $userRepository;

    public function __construct(ReserveRepositoryInterface $reserveRepository, TableRepositoryInterface $tableRepository, UserRepositoryInterface $userRepository)
    {
        $this->reserveRepository = $reserveRepository;
        $this->tableRepository = $tableRepository;
        $this->userRepository = $userRepository;
    }

    public function getAllReservesService()
    {
        return ReserveResource::collection($this->reserveRepository->getAllReservesRepo());
    }

    public function getReserveByIdService($id)
    {
        try {
            return ReserveResource::make($this->reserveRepository->getReserveByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Reserve doesn't exist"
            ]);
        }
    }

    public function createReserveService(CreateReserve $request)
    {
        try {
            UserResource::make($this->userRepository->getUserByIdRepo($request->get('user')));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "User doesn't exist"
            ]);
        }

        try {
            TableResource::make($this->tableRepository->getTableByIdRepo($request->get('table')));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Table doesn't exist"
            ]);
        }

        try {
            $date = Carbon::createFromFormat('d/m/Y', $request->get('date'))->startOfDay();
            if ($date->isBefore(Carbon::now()->startOfDay())) {
                return response()->json([
                    "Message" => "You can't reserve this date"
                ]);
            }
            return ReserveResource::make($this->reserveRepository->createReserveRepo($request));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Reserve already exists"
            ]);
        }
    }

    public function updateReserveService(UpdateReserve $request, $id)
    {
        try {
            ReserveResource::make($this->reserveRepository->getReserveByIdRepo($id));
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Reserve doesn't exist"
            ]);
        }

        try {
            $update = $this->reserveRepository->updateReserveRepo($request, $id);
            if ($update == 1) {
                return ReserveResource::make($this->reserveRepository->getReserveByIdRepo($id));
            } else {
                return response()->json([
                    "Message" => "Not found"
                ], 404);
            };
        } catch (\Exception $e) {
            return response()->json([
                "Message" => "Update error"
            ]);
        }
    }

    public function deleteReserveService($id)
    {
        $delete = $this->reserveRepository->deleteReserveRepo($id);

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
