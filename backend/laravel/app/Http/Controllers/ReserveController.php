<?php

namespace App\Http\Controllers;

use App\Http\Requests\Reserve\CreateReserve;
use App\Http\Requests\Reserve\UpdateReserve;
use App\Models\Reserve;
use App\Services\ReserveService;

use Illuminate\Http\JsonResponse;

class ReserveController extends Controller
{
    private $reserveService;

    public function __construct(ReserveService $reserveService)
    {
        $this->reserveService = $reserveService;
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index()
    {
        return $this->reserveService->getAllReservesService();
    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \App\Http\Requests\Reserve\CreateReserve  $request
     * @return \Illuminate\Http\Response
     */
    public function store(CreateReserve $request)
    {
        return $this->reserveService->createReserveService($request);
    }

    /**
     * Display the specified resource.
     *
     * @param  \App\Models\Reserve  $reserve
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        return $this->reserveService->getReserveByIdService($id);
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  \App\Models\Reserve  $reserve
     * @return \Illuminate\Http\Response
     */
    public function edit(Reserve $reserve)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \App\Http\Requests\Reserve\UpdateReserve $request
     * @param  \App\Models\Reserve  $reserve
     * @return \Illuminate\Http\Response
     */
    public function update(UpdateReserve $request, $id)
    {
        return $this->reserveService->updateReserveService($request,$id);
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  \App\Models\Reserve  $reserve
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
        return $this->reserveService->deleteReserveService($id);
    }
}
