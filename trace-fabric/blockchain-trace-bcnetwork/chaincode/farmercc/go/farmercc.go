package main
import (
        "bytes"
        "encoding/json"
        "fmt"
       // "strconv"
        "github.com/hyperledger/fabric/core/chaincode/shim"
        sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {
}

type Crops struct {
	//操作ID
        CropsId   string `json:"crops_id"`
        //操作员(语音转文本)姓名
        CropsName  string `json:"crops_name"`
	//地址
	Address string `json:"address"`
	//审查时间
	RegisterTime string `json:"register_time"`
        //通话时长
	Year string `json:"year"`
	//通话人员1
        FarmerName string `json:"farmer_name"`
	//通话人员1 ID
	FarmerID string `json:"farmer_id"`
        //通话人员1电话
	FarmerTel string `json:"farmer_tel"`
	//通话人员2
	FertilizerName string `json:"fertilizer_name"`
	//通话方式
	PlatMode string `json:"plant_mode"`
	//通话主题
	BaggingStatus string `json:"bagging_status"`
	//通话人员2电话
	GrowSeedlingsCycle string `json:"grow_seedlings_cycle"`
	//原始语音哈希
	IrrigationCycle string `json:"irrigation_cycle"`
	//施肥周期
	ApplyFertilizerCycle string `json:"apply_fertilizer_cycle"`
	//审查结果
	WeedCycle string `json:"weed_cycle"`
	//溯源码
	Remarks string `json:"remarks"`

}

type CropsGrowInfo struct {
        //生成情况唯一ID
	CropsGrowId string `json:"crops_grow_id"`
        //ID
        CropsBakId string `json:"crops_bak_id"`
        //记录时间
        RecordTime string `json:"record_time"`
        //图片URL
        CropsGrowPhotoUrl string `json:"crops_grow_photo_url"`
        //温度
        Temperature string `json:"temperature"`
	//设备情况
	GrowStatus string `json:"grow_status"`
	//湿度
	WaterContent string `json:"water_content"`
	//光线情况
	IlluminationStatus string `json:"illumination_status"`
        //溯源码
        Remarks string `json:"remarks"`

}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
        return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {
        function, args := APIstub.GetFunctionAndParameters()
        // Route to the appropriate handler function to interact with the ledger appropriately
		//根据id查询
        if function == "queryCropsById" {
             return s.queryCropsById(APIstub, args)
        }else if function == "initLedger" {
             return s.initLedger(APIstub)
        }else if function == "createCrops" {
             return s.createCrops(APIstub, args)
        }else if function == "queryCropsProcessByCropsId" {
             return s.queryCropsProcessByCropsId(APIstub,args)
        }else if function == "recordCropsGrow" {
             return s.recordCropsGrow(APIstub, args)
        }
        return shim.Error("Invalid Smart Contract function name.")
}

/**
 * 根据ID查询信息
 */
func (s *SmartContract) queryCropsById(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting 1")
        }
        cropsAsBytes, _ := APIstub.GetState(args[0])
        return shim.Success(cropsAsBytes)
}

//根据crops_id溯源所有记录过程
func (s *SmartContract) queryCropsProcessByCropsId(APIstub shim.ChaincodeStubInterface,args[]string) sc.Response{
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	CropsBakId := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"crops_bak_id\":{\"$eq\":\"%s\"}}}", CropsBakId)
	resultsIterator, err := APIstub.GetQueryResult(queryString)
        if err != nil {
                return shim.Error(err.Error())
        }
        defer resultsIterator.Close()

        // buffer is a JSON array containing QueryResults
        var buffer bytes.Buffer
        buffer.WriteString("[")

        bArrayMemberAlreadyWritten := false
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
                        return shim.Error(err.Error())
                }
                // Add a comma before array members, suppress it for the first array member
                if bArrayMemberAlreadyWritten == true {
                        buffer.WriteString(",")
                }
                buffer.WriteString("{\"Key\":")
                buffer.WriteString("\"")
                buffer.WriteString(queryResponse.Key)
                buffer.WriteString("\"")

                buffer.WriteString(", \"Record\":")
                // Record is a JSON object, so we write as-is
                buffer.WriteString(string(queryResponse.Value))
                buffer.WriteString("}")
                bArrayMemberAlreadyWritten = true
        }
        buffer.WriteString("]")

        fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

        return shim.Success(buffer.Bytes())
}


/**
 * 初始化账本
 */
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
        crops := []Crops{
		 Crops{CropsId: "liu_first_record",
                 CropsName: "LiuZhen",
                 Address: "XuZhou",
                 RegisterTime: "2024.09.21",
		         Year: "20min",
                 FarmerName: "LiUPAI",
                 FarmerID: "342319492349",
                 FertilizerName: "ZhangShuai",
                 PlatMode: "对讲机",
                 BaggingStatus: "1234",
                 GrowSeedlingsCycle: "13897344234",
                 IrrigationCycle: "2b4be5137612d4fe2546d315a78750d7cfe37a959d8b10291b25d301fa296a1f",
                 ApplyFertilizerCycle: "语言文本验证通过，无异常,生产符合安全标准.",
                 WeedCycle: "车皮数量问题",
                 Remarks: "848768919926476800"},
                        }
        i := 0
        for i < len(crops) {
                fmt.Println("i is ", i)
                cropsAsBytes, _ := json.Marshal(crops[i])
                APIstub.PutState(crops[i].CropsId, cropsAsBytes)
                fmt.Println("Added", crops[i])
                i = i + 1
        }
        return shim.Success(nil)
}

/**
 *记录作物生长
 */
func (s *SmartContract) recordCropsGrow(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) !=  10{
                return shim.Error("Incorrect number of arguments. Expecting 10")
        }
        var cropsGrowInfo = CropsGrowInfo{CropsGrowId: args[1], CropsBakId: args[2], RecordTime: args[3], CropsGrowPhotoUrl: args[4], Temperature: args[5], GrowStatus: args[6], WaterContent: args[7], IlluminationStatus: args[8], Remarks: args[9]}
        cropsGrowInfoAsBytes, _ := json.Marshal(cropsGrowInfo)
        APIstub.PutState(args[0],cropsGrowInfoAsBytes)
        return shim.Success(nil)
}

/**
 *添加作物
 */
func (s *SmartContract) createCrops(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 if len(args) != 16 {
	        return shim.Error("Incorrect number of arguments. Expecting 15")
	 }
         var crops = Crops{CropsId: args[1], CropsName: args[2], Address: args[3],FarmerName: args[4],FertilizerName: args[5], PlatMode: args[6],BaggingStatus: args[7],GrowSeedlingsCycle: args[8],IrrigationCycle: args[9],ApplyFertilizerCycle: args[10],WeedCycle: args[11],Remarks: args[12],RegisterTime: args[13],Year: args[14],FarmerTel: args[15]}
	 cropsAsBytes, _ := json.Marshal(crops)
         APIstub.PutState(args[0], cropsAsBytes)
         return shim.Success(nil)
}

func (s *SmartContract) queryAllCropsByFarmerID(APIstub shim.ChaincodeStubInterface,args string) sc.Response {


        startKey := "CAR0"
        endKey := "CAR999"

        resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
        if err != nil {
                return shim.Error(err.Error())
        }
        defer resultsIterator.Close()

        // buffer is a JSON array containing QueryResults
        var buffer bytes.Buffer
        buffer.WriteString("[")

        bArrayMemberAlreadyWritten := false
        for resultsIterator.HasNext() {
                queryResponse, err := resultsIterator.Next()
                if err != nil {
                        return shim.Error(err.Error())
                }
                // Add a comma before array members, suppress it for the first array member
                if bArrayMemberAlreadyWritten == true {
                        buffer.WriteString(",")
                }
                buffer.WriteString("{\"Key\":")
                buffer.WriteString("\"")
                buffer.WriteString(queryResponse.Key)
                buffer.WriteString("\"")

                buffer.WriteString(", \"Record\":")
                // Record is a JSON object, so we write as-is
                buffer.WriteString(string(queryResponse.Value))
                buffer.WriteString("}")
                bArrayMemberAlreadyWritten = true
        }
        buffer.WriteString("]")

        fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

        return shim.Success(buffer.Bytes())
}


func main() {
        err := shim.Start(new(SmartContract))
        if err != nil {
                fmt.Printf("Error creating new Smart Contract: %s", err)
        }

}

