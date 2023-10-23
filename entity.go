package wujiesdk

// @Title        entity.go
// @Description  entity
// @Create       XdpCs 2023-09-10 20:47
// @Update       XdpCs 2023-10-23 15:00

import (
	"fmt"
)

type BaseResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type AvailableIntegralBalanceResponse struct {
	BaseResponse
	Data struct {
		Balance int `json:"balance"`
	} `json:"data"`
}

type ExchangePointRequest struct {
	ExchangeTargetMobile string `json:"exchange_target_mobile"`
	Amount               int    `json:"amount"`
}

func (e *ExchangePointRequest) String() string {
	return fmt.Sprintf("%+v", *e)
}

type ModelBaseInfosResponse struct {
	BaseResponse
	Data []ModelBaseInfoData `json:"data"`
}

type ModelBaseInfoData struct {
	ModelType         int32  `json:"type"`
	ModelCode         int32  `json:"model_code"`
	ModelVersion      string `json:"model_version"`
	ModelDesc         string `json:"model_desc"`
	ControlNetSupport string `json:"controlnet_support"`
}

type DefaultResourceStyleModelResponse struct {
	BaseResponse
	Data DefaultResourceStyleModelData `json:"data"`
}

type DefaultResourceStyleModelData struct {
	StyleModels []StyleModel `json:"style_model"`
}

type StyleModel struct {
	Key            string `json:"key"`
	Name           string `json:"name"`
	ModelCode      int    `json:"model_code"`
	SampleImageURL string `json:"sample_image_url"`
}

type DefaultResourceModelResponse struct {
	BaseResponse
	Data DefaultResourceModelData `json:"data"`
}

type DefaultResourceModelData struct {
	CreateOptionMenu struct {
		ImageType []struct {
			Name     string `json:"name"`
			Url      string `json:"url"`
			Category string `json:"category"`
		} `json:"image_type"`
		PromptTips []struct {
			Name string `json:"name"`
		} `json:"prompt_tips"`
		Resolution []struct {
			Width              int   `json:"width"`
			Height             int   `json:"height"`
			SuperSizeMultiple  int   `json:"super_size_multiple"`
			PrefineMultiples   []int `json:"prefine_multiples"`
			SuperSizeMultiples []int `json:"super_size_multiples"`
			SuperSizeDetails   []struct {
				Multiple      float64 `json:"multiple"`
				IntegralPrice int     `json:"integral_price"`
			} `json:"super_size_details"`
			PrefineDetails []struct {
				Multiple      float64 `json:"multiple"`
				IntegralPrice int     `json:"integral_price"`
			} `json:"prefine_details"`
			Url       string `json:"url"`
			SizeRatio string `json:"size_ratio"`
		} `json:"resolution"`
		ResolutionNew struct {
			ResolutionKey  string `json:"resolution_key"`
			ResolutionList []struct {
				Width             int     `json:"width"`
				Height            int     `json:"height"`
				SuperSizeMultiple int     `json:"super_size_multiple"`
				PrefineMultiples  float64 `json:"prefine_multiples"`
				DisplayResolution string  `json:"display_resolution"`
				Url               string  `json:"url"`
				SizeRatio         string  `json:"size_ratio"`
			} `json:"resolution_list"`
		} `json:"resolution_new"`
		Style []struct {
			Name     string `json:"name"`
			Url      string `json:"url"`
			Category string `json:"category"`
		} `json:"style"`
		Artist []struct {
			Name     string `json:"name"`
			Url      string `json:"url"`
			Category string `json:"category"`
		} `json:"artist"`
		ElementMagic []struct {
			Key       string `json:"key"`
			Name      string `json:"name"`
			ChoiceKey string `json:"choice_key"`
		} `json:"element_magic"`
		StyleDecoration []struct {
			Key       string `json:"key"`
			Name      string `json:"name"`
			ChoiceKey string `json:"choice_key"`
		} `json:"style_decoration"`
		Character []struct {
			Key                  string   `json:"key"`
			Name                 string   `json:"name"`
			Category             string   `json:"category"`
			RecommendedWeight    int      `json:"recommended_weight"`
			SupportModelVersions []string `json:"support_model_versions"`
		} `json:"character"`
		ModelFusion []struct {
			Key                  string   `json:"key"`
			Name                 string   `json:"name"`
			Category             string   `json:"category"`
			RecommendedWeight    int      `json:"recommended_weight"`
			SupportModelVersions []string `json:"support_model_versions"`
		} `json:"model_fusion"`
		Patterns []struct {
			Name string `json:"name"`
		} `json:"patterns"`
		SamplerModels []struct {
			SamplerModelName string `json:"sampler_model_name"`
			SamplerIndex     int    `json:"sampler_index"`
		} `json:"sampler_models"`
	} `json:"create_option_menu"`
}

type CreateImageRequest struct {
	Model               int      `json:"model"`
	Prompt              string   `json:"prompt"`
	UcPrompt            string   `json:"uc_prompt"`
	FullyCustomUcPrompt bool     `json:"fully_custom_uc_prompt"`
	Num                 int      `json:"num"`
	Width               int      `json:"width"`
	Height              int      `json:"height"`
	InitImageURL        string   `json:"init_image_url"`
	InitWidth           int      `json:"init_width"`
	InitHeight          int      `json:"init_height"`
	CreativityDegree    int      `json:"creativity_degree"`
	InitImageSimilarity int      `json:"init_image_similarity"`
	SuperSizeMultiple   int      `json:"super_size_multiple"`
	PrefineMultiple     int      `json:"prefine_multiple"`
	ImageType           []string `json:"image_type"`
	Style               []string `json:"style"`
	Artist              string   `json:"artist"`
	Artists             []string `json:"artists"`
	ElementMagic        []string `json:"element_magic"`
	StyleDecoration     []string `json:"style_decoration"`
	ModelParam          string   `json:"model_param"`
	AccelerateTimes     int      `json:"accelerate_times"`
	Vendor              int      `json:"vendor"`
	Character           []string `json:"character"`
	ModelFusion         []struct {
		Key    string  `json:"key"`
		Weight float64 `json:"weight"`
	} `json:"model_fusion"`
	StyleModel         string `json:"style_model"`
	Pattern            string `json:"pattern"`
	PretreatmentMethod string `json:"pretreatment_method"`
	Steps              int    `json:"steps"`
	Cfg                int    `json:"cfg"`
	SamplerIndex       int    `json:"sampler_index"`
	Seed               string `json:"seed"`
	QueueType          int    `json:"queue_type"`
	CreateSource       int    `json:"create_source"`
	ClipSkip           int    `json:"clip_skip"`
	ControlWeight      int    `json:"control_weight"`
	ControlImg2Img     bool   `json:"control_img2_img"`
	ControlMode        int    `json:"control_mode"`
	DetectGrayNum      int    `json:"detect_gray_num"`
	NotifyURL          string `json:"notify_url"`
	ServiceContext     struct {
		Source         string `json:"source"`
		From           string `json:"from"`
		UserAgent      string `json:"user_agent"`
		AppCode        string `json:"app_code"`
		DeviceId       string `json:"device_id"`
		Ip             string `json:"ip"`
		RegisterSource string `json:"register_source"`
	} `json:"service_context"`
	MultiDiffusion struct {
		TiledDiffusion struct {
			Enabled                     bool   `json:"enabled"`
			Method                      string `json:"method"`
			OverwriteSize               bool   `json:"overwrite_size"`
			KeepInputSize               bool   `json:"keep_input_size"`
			ImageWidth                  int    `json:"image_width"`
			ImageHeight                 int    `json:"image_height"`
			TileWidth                   int    `json:"tile_width"`
			TileHeight                  int    `json:"tile_height"`
			Overlap                     int    `json:"overlap"`
			TileBatchSize               int    `json:"tile_batch_size"`
			UpscalerName                string `json:"upscaler_name"`
			ScaleFactor                 int    `json:"scale_factor"`
			NoiseInverse                bool   `json:"noise_inverse"`
			NoiseInverseSteps           int    `json:"noise_inverse_steps"`
			NoiseInverseRetouch         int    `json:"noise_inverse_retouch"`
			NoiseInverseRenoiseStrength int    `json:"noise_inverse_renoise_strength"`
			NoiseInverseRenoiseKernel   int    `json:"noise_inverse_renoise_kernel"`
			ControlTensorCpu            bool   `json:"control_tensor_cpu"`
			EnableBboxControl           bool   `json:"enable_bbox_control"`
			DrawBackground              bool   `json:"draw_background"`
			CausalLayers                bool   `json:"causal_layers"`
			BboxControlStates           []struct {
				Enabled      bool    `json:"enabled"`
				X            float64 `json:"x"`
				Y            float64 `json:"y"`
				W            float64 `json:"w"`
				H            float64 `json:"h"`
				Prompt       string  `json:"prompt"`
				NegPrompt    string  `json:"neg_prompt"`
				BlendMode    string  `json:"blend_mode"`
				FeatherRatio float64 `json:"feather_ratio"`
				Seed         int     `json:"seed"`
			} `json:"bbox_control_states"`
		} `json:"tiled_diffusion"`
		TiledVae struct {
			Enabled         bool `json:"enabled"`
			EncoderTileSize int  `json:"encoder_tile_size"`
			DecoderTileSize int  `json:"decoder_tile_size"`
			VaeToGpu        bool `json:"vae_to_gpu"`
			FastDecoder     bool `json:"fast_decoder"`
			FastEncoder     bool `json:"fast_encoder"`
			ColorFix        bool `json:"color_fix"`
		} `json:"tiled_vae"`
	} `json:"multi_diffusion"`
	ServiceType int    `json:"service_type"`
	Extra       string `json:"extra"`
	ProMethod   string `json:"pro_method"`
}

func (c *CreateImageRequest) String() string {
	return fmt.Sprintf("%+v", *c)
}

type CreateImageResponse struct {
	BaseResponse
	Data CreateImageData `json:"data"`
}

type CreateImageData struct {
	Keys    []string `json:"keys"`
	Results []struct {
		Key            string `json:"key"`
		ExpectedSecond int    `json:"expected_second"`
	} `json:"results"`
	ExpectedIntegralCost int `json:"expected_integral_cost"`
}

type GeneratingInfoResponse struct {
	BaseResponse
	Data GeneratingInfoData `json:"data"`
}

type GeneratingInfoData struct {
	List []ImageGeneratingInfo `json:"list"`
}

type ImageGeneratingInfo struct {
	Key             string  `json:"key"`
	Status          int     `json:"status"`
	PictureURL      string  `json:"picture_url"`
	ExpectedSeconds int     `json:"expected_seconds"`
	StartGenTime    int     `json:"start_gen_time"`
	CompleteTime    int     `json:"complete_time"`
	CompletePercent float64 `json:"complete_percent"`
	QueueBeforeNum  int     `json:"queue_before_num"`
	ReduceTime      int     `json:"reduce_time"`
	InvolveYellow   int     `json:"involve_yellow"`
	AuditInfo       string  `json:"audit_info"`
	FailMessage     struct {
		FailCode    int    `json:"fail_code"`
		FailMessage string `json:"fail_message"`
	} `json:"fail_message"`
	ModelPrompt         string `json:"model_prompt"`
	IntegralCost        int    `json:"integral_cost"`
	IntegralCostMessage string `json:"integral_cost_message"`
}

type ImageInfoResponse struct {
	BaseResponse
	Data ImageInfoData `json:"data"`
}

type ImageInfoData struct {
	Prompt              string   `json:"prompt"`
	UcPrompt            string   `json:"uc_prompt"`
	Model               int      `json:"model"`
	Width               int      `json:"width"`
	Height              int      `json:"height"`
	Status              int      `json:"status"`
	PictureUrl          string   `json:"picture_url"`
	MiniPictureURL      string   `json:"mini_picture_url"`
	InitImageURL        string   `json:"init_image_url"`
	InitImageSimilarity int      `json:"init_image_similarity"`
	CreativityDegree    int      `json:"creativity_degree"`
	Artist              string   `json:"artist"`
	Style               string   `json:"style"`
	ImageType           string   `json:"image_type"`
	ElementMagic        []string `json:"element_magic"`
	GenerateTime        int      `json:"generate_time"`
	StartGenTime        int      `json:"start_gen_time"`
	CompleteTime        int      `json:"complete_time"`
	InvolveYellow       int      `json:"involve_yellow"`
	AuditInfo           string   `json:"audit_info"`
	TechnologyInfo      struct {
		MachineNo        string  `json:"machine_no"`
		GpuType          string  `json:"gpu_type"`
		PowerConsumption float64 `json:"power_consumption"`
	} `json:"technology_info"`
	FailMessage struct {
		FailCode    int    `json:"fail_code"`
		FailMessage string `json:"fail_message"`
	} `json:"fail_message"`
	ModelPrompt      string   `json:"model_prompt"`
	CharacterOptions []string `json:"character_options"`
	ModelFusion      []struct {
		Name   string  `json:"name"`
		Weight float64 `json:"weight"`
	} `json:"model_fusion"`
	StyleModel          string  `json:"style_model"`
	PretreatmentImage   string  `json:"pretreatment_image"`
	PretreatmentMethod  string  `json:"pretreatment_method"`
	Steps               int     `json:"steps"`
	Cfg                 float64 `json:"cfg"`
	SamplerIndex        int     `json:"sampler_index"`
	Seed                string  `json:"seed"`
	IntegralCost        int     `json:"integral_cost"`
	IntegralCostMessage string  `json:"integral_cost_message"`
	MultiDiffusion      struct {
		TiledDiffusion struct {
			Enabled                     bool   `json:"enabled"`
			Method                      string `json:"method"`
			OverwriteSize               bool   `json:"overwrite_size"`
			KeepInputSize               bool   `json:"keep_input_size"`
			ImageWidth                  int    `json:"image_width"`
			ImageHeight                 int    `json:"image_height"`
			TileWidth                   int    `json:"tile_width"`
			TileHeight                  int    `json:"tile_height"`
			Overlap                     int    `json:"overlap"`
			TileBatchSize               int    `json:"tile_batch_size"`
			UpscalerName                string `json:"upscaler_name"`
			ScaleFactor                 int    `json:"scale_factor"`
			NoiseInverse                bool   `json:"noise_inverse"`
			NoiseInverseSteps           int    `json:"noise_inverse_steps"`
			NoiseInverseRetouch         int    `json:"noise_inverse_retouch"`
			NoiseInverseRenoiseStrength int    `json:"noise_inverse_renoise_strength"`
			NoiseInverseRenoiseKernel   int    `json:"noise_inverse_renoise_kernel"`
			ControlTensorCpu            bool   `json:"control_tensor_cpu"`
			EnableBboxControl           bool   `json:"enable_bbox_control"`
			DrawBackground              bool   `json:"draw_background"`
			CausalLayers                bool   `json:"causal_layers"`
			BboxControlStates           []struct {
				Enabled      bool    `json:"enabled"`
				X            float64 `json:"x"`
				Y            float64 `json:"y"`
				W            float64 `json:"w"`
				H            float64 `json:"h"`
				Prompt       string  `json:"prompt"`
				NegPrompt    string  `json:"neg_prompt"`
				BlendMode    string  `json:"blend_mode"`
				FeatherRatio float64 `json:"feather_ratio"`
				Seed         int     `json:"seed"`
			} `json:"bbox_control_states"`
		} `json:"tiled_diffusion"`
		TiledVae struct {
			Enabled         bool `json:"enabled"`
			EncoderTileSize int  `json:"encoder_tile_size"`
			DecoderTileSize int  `json:"decoder_tile_size"`
			VaeToGpu        bool `json:"vae_to_gpu"`
			FastDecoder     bool `json:"fast_decoder"`
			FastEncoder     bool `json:"fast_encoder"`
			ColorFix        bool `json:"color_fix"`
		} `json:"tiled_vae"`
	} `json:"multi_diffusion"`
}

type ImagePriceInfoRequest struct {
	CreateImageRequest
}

func (i *ImagePriceInfoRequest) String() string {
	return fmt.Sprintf("%+v", *i)
}

type ImagePriceInfoResponse struct {
	BaseResponse
	Data ImagePriceInfoData `json:"data"`
}

type ImagePriceInfoData struct {
	UserRightsUse struct {
		DailyFreeCreateTimesUse int `json:"daily_free_create_times_use"`
		DeservedCreateTimesUse  int `json:"deserved_create_times_use"`
	} `json:"user_rights_use"`
	VipRightsUse struct {
		CreateTimesUse     int `json:"create_times_use"`
		AccelerateTimesUse int `json:"accelerate_times_use"`
		SuperSizeTimesUse  int `json:"super_size_times_use"`
	} `json:"vip_rights_use"`
	IntegralUse struct {
		IntegralUseOnCreate     int `json:"integral_use_on_create"`
		IntegralUseOnResolution int `json:"integral_use_on_resolution"`
		IntegralUseOnStyleModel int `json:"integral_use_on_style_model"`
		IntegralUseOnSteps      int `json:"integral_use_on_steps"`
		IntegralUseOnAccelerate int `json:"integral_use_on_accelerate"`
		IntegralUseOnSuperSize  int `json:"integral_use_on_super_size"`
		DiscountIntegral        int `json:"discount_integral"`
	} `json:"integral_use"`
}

type PostSuperSizeRequest struct {
	URL           string            `json:"url"`
	Multiple      float64           `json:"multiple"`
	SuperSizeType SuperSizeType     `json:"super_size_type"`
	CostType      SuperSizeCostType `json:"cost_type"`
}

type SuperSizeOption func(s *PostSuperSizeRequest)

func NewPostSuperSizeRequest(url string, multiple float64, options ...SuperSizeOption) *PostSuperSizeRequest {
	p := &PostSuperSizeRequest{URL: url, Multiple: multiple}
	for _, option := range options {
		option(p)
	}
	return p
}

func WithSuperSizeType(superSizeType SuperSizeType) SuperSizeOption {
	return func(p *PostSuperSizeRequest) {
		p.SuperSizeType = superSizeType
	}
}

func WithCostType(costType SuperSizeCostType) func(s *PostSuperSizeRequest) {
	return func(p *PostSuperSizeRequest) {
		p.CostType = costType
	}
}

func (s *PostSuperSizeRequest) String() string {
	return fmt.Sprintf("%+v", *s)
}

type PostSuperSizeResponse struct {
	BaseResponse
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}

type GetSuperSizeResponse struct {
	BaseResponse
	Data []SuperSizeInfo `json:"data"`
}

type SuperSizeInfo struct {
	Key      string  `json:"key"`
	URL      string  `json:"url"`
	SrURL    string  `json:"sr_url"`
	Multiple float64 `json:"multiple"`
	Status   int     `json:"status"`
	Integral int     `json:"integral"`
	Duration int     `json:"duration"`
}

type CreateParamsResponse struct {
	BaseResponse
	Data []CreateParams `json:"data"`
}

type CreateParams struct {
	Key                      string  `json:"key"`
	ArtworkURL               string  `json:"artwork_url"`
	Model                    int     `json:"model"`
	ModelAsString            string  `json:"model_as_string"`
	ModelCode                int     `json:"model_code"`
	ModelCodeAsString        string  `json:"model_code_as_string"`
	Pattern                  string  `json:"pattern"`
	Prompt                   string  `json:"prompt"`
	UcPrompt                 string  `json:"uc_prompt"`
	CreativityDegree         int     `json:"creativity_degree"`
	CreativityDegreeAsString string  `json:"creativity_degree_as_string"`
	InitImageURL             string  `json:"init_image_url"`
	InitWidth                int     `json:"init_width"`
	InitWidthAsString        string  `json:"init_width_as_string"`
	InitHeight               int     `json:"init_height"`
	InitHeightAsString       string  `json:"init_height_as_string"`
	PretreatmentMethod       string  `json:"pretreatment_method"`
	MaskImageURL             string  `json:"mask_image_url"`
	MaskZoneImageURL         string  `json:"mask_zone_image_url"`
	Size                     string  `json:"size"`
	Nature                   int     `json:"nature"`
	NatureAsString           string  `json:"nature_as_string"`
	PromptOptimize           int     `json:"prompt_optimize"`
	PromptOptimizeAsString   string  `json:"prompt_optimize_as_string"`
	StyleDecoration          string  `json:"style_decoration"`
	Character                string  `json:"character"`
	ModelFusion              string  `json:"model_fusion"`
	StyleModel               string  `json:"style_model"`
	ResolutionInfo           string  `json:"resolution_info"`
	Steps                    int     `json:"steps"`
	StepsAsString            string  `json:"steps_as_string"`
	Cfg                      float64 `json:"cfg"`
	CfgAsString              string  `json:"cfg_as_string"`
	SamplerIndex             string  `json:"sampler_index"`
	Seed                     string  `json:"seed"`
	SuperType                int     `json:"super_type"`
	SuperTypeAsString        string  `json:"super_type_as_string"`
	ChatGptOptimize          bool    `json:"chat_gpt_optimize"`
	ChatGptOptimizeAsString  string  `json:"chat_gpt_optimize_as_string"`
	ClipSkip                 int     `json:"clip_skip"`
	ClipSkipAsString         string  `json:"clip_skip_as_string"`
	Ensd                     float64 `json:"ensd"`
	EnsdAsString             string  `json:"ensd_as_string"`
	RepairTheHand            bool    `json:"repair_the_hand"`
	RepairTheHandAsString    string  `json:"repair_the_hand_as_string"`
	ConsumedTime             string  `json:"consumed_time"`
}

type ImageModelQueueInfoResponse struct {
	BaseResponse
	Data ImageModelQueueInfoData `json:"data"`
}

type ImageModelQueueInfoData struct {
	ExpectedSeconds int `json:"expected_seconds"`
	QueueNum        int `json:"queue_num"`
	ReduceTime      int `json:"reduce_time"`
}

type CancelImageResponse struct {
	BaseResponse
	Data string `json:"data"`
}

type AccelerateImageRequest struct {
	Key     string `json:"key"`
	StepNum int    `json:"step_num"`
}

func NewAccelerateImageRequest(key string, stepNum int) *AccelerateImageRequest {
	return &AccelerateImageRequest{Key: key, StepNum: stepNum}
}

func (a *AccelerateImageRequest) String() string {
	return fmt.Sprintf("%+v", *a)
}

type AccelerateImageResponse struct {
	BaseResponse
	Data string `json:"data"`
}

type PromptOptimizeSubmitRequest struct {
	TaskID      string               `json:"task_id"`
	Type        PromptSubmitType     `json:"type"`
	Original    string               `json:"original"`
	Language    PromptSubmitLanguage `json:"language"`
	CallbackURL string               `json:"callback_url"`
}

func (p *PromptOptimizeSubmitRequest) String() string {
	return fmt.Sprintf("%+v", *p)
}

type PromptOptimizeSubmitResponse struct {
	BaseResponse
}

type PromptOptimizeResultResponse struct {
	BaseResponse
	Data PromptOptimizeResultData `json:"data"`
}

type PromptOptimizeResultData struct {
	TaskID string `json:"task_id"`
	Code   int    `json:"code"`
	Result string `json:"result"`
}

type YouthifyRequest struct {
	ImageURL          string  `json:"image_url"`
	InitWidth         int     `json:"init_width"`
	InitHeight        int     `json:"init_height"`
	Width             int     `json:"width"`
	Height            int     `json:"height"`
	SuperSizeMultiple float64 `json:"super_size_multiple"`
	NotifyURL         string  `json:"notify_url"`
}

func (y *YouthifyRequest) String() string {
	return fmt.Sprintf("%+v", *y)
}

type YouthifyResponse struct {
	BaseResponse
	Data YouthifyData `json:"data"`
}

type YouthifyData struct {
	Keys    string `json:"keys"`
	Results []struct {
		Key            string `json:"key"`
		ExpectedSecond int    `json:"expected_second"`
	} `json:"results"`
	ExpectedIntegralCost int `json:"expected_integral_cost"`
}

type QuerySpellResponse struct {
	BaseResponse
	Data []QuerySpellData `json:"data"`
}

type QuerySpellData struct {
	SpellName   string `json:"spell_name"`
	SpellEnName string `json:"spell_en_name"`
	Icon        string `json:"icon"`
	Category    string `json:"category"`
	Label       string `json:"label"`
}

type CreateImageProRequest struct {
	ModelCode         int    `json:"model_code"`
	Prompt            string `json:"prompt"`
	Width             int    `json:"width"`
	Height            int    `json:"height"`
	SupersizeMultiple int    `json:"supersize_multiple"`
	PrefineMultiple   int    `json:"prefine_multiple"`
	BatchCount        int    `json:"batch_count"`
	OptionParam       struct {
		ModelFusion []struct {
			Key    string  `json:"key"`
			Weight float64 `json:"weight"`
		} `json:"model_fusion"`
		Character []string `json:"character"`
	} `json:"option_param"`
	AdvancedParam struct {
		UcPrompt     string  `json:"uc_prompt"`
		RestoreFaces bool    `json:"restore_faces"`
		Tilling      bool    `json:"tilling"`
		Seed         string  `json:"seed"`
		VaeFile      string  `json:"vae_file"`
		Cfg          int     `json:"cfg"`
		SamplerSteps int     `json:"sampler_steps"`
		SamplerIndex int     `json:"sampler_index"`
		ClipSkip     int     `json:"clip_skip"`
		Ensd         float64 `json:"ensd"`
	} `json:"advanced_param"`
	ImgToImgParam struct {
		InitImageUrl     string `json:"init_image_url"`
		CreativityDegree int    `json:"creativity_degree"`
		ResizeMode       int    `json:"resize_mode"`
	} `json:"img_to_img_param"`
	ControlNetParams []struct {
		Type                int    `json:"type"`
		Preprocessor        int    `json:"preprocessor"`
		Model               int    `json:"model"`
		ControlWeight       int    `json:"control_weight"`
		StartingControlStep int    `json:"starting_control_step"`
		EndingControlStep   int    `json:"ending_control_step"`
		ControlMode         int    `json:"control_mode"`
		ImageUrl            string `json:"image_url"`
		ImageWidth          int    `json:"image_width"`
		ImageHeight         int    `json:"image_height"`
		Mask                string `json:"mask"`
		MaskUrl             string `json:"mask_url"`
		ProcessorRes        int    `json:"processor_res"`
		ThresholdA          int    `json:"threshold_a"`
		ThresholdB          int    `json:"threshold_b"`
		ResizeMode          int    `json:"resize_mode"`
		PixelPerfect        bool   `json:"pixel_perfect"`
	} `json:"control_net_params"`
	InpaintingPluginDTO struct {
		MaskZoneImageUrl      string `json:"mask_zone_image_url"`
		MaskBlur              int    `json:"mask_blur"`
		InpaintingFill        int    `json:"inpainting_fill"`
		InpaintingMaskInvert  bool   `json:"inpainting_mask_invert"`
		InpaintFullResPadding int    `json:"inpaint_full_res_padding"`
		InpaintFullRes        bool   `json:"inpaint_full_res"`
	} `json:"inpainting_plugin_d_t_o"`
	TiledDiffusionDTO struct {
		Enabled           bool `json:"enabled"`
		DrawBackground    bool `json:"draw_background"`
		BboxControlStates []struct {
			Enabled             bool    `json:"enabled"`
			X                   float64 `json:"x"`
			Y                   float64 `json:"y"`
			W                   float64 `json:"w"`
			H                   float64 `json:"h"`
			Prompt              string  `json:"prompt"`
			NegPrompt           string  `json:"neg_prompt"`
			ModelInputPrompt    string  `json:"model_input_prompt"`
			ModelInputNegPrompt string  `json:"model_input_neg_prompt"`
			BlendMode           string  `json:"blend_mode"`
			Seed                int     `json:"seed"`
			OptionParam         struct {
				ModelFusion []struct {
					Key    string  `json:"key"`
					Weight float64 `json:"weight"`
				} `json:"model_fusion"`
				Character []string `json:"character"`
			} `json:"option_param"`
		} `json:"bbox_control_states"`
	} `json:"tiled_diffusion_d_t_o"`
	FaceEditorDTO struct {
		Enabled                 bool     `json:"enabled"`
		UseMinimalArea          bool     `json:"use_minimal_area"`
		AffectedAreas           []string `json:"affected_areas"`
		MaskSize                int      `json:"mask_size"`
		MaskBlur                int      `json:"mask_blur"`
		MaxFaceCount            int      `json:"max_face_count"`
		Confidence              float64  `json:"confidence"`
		FaceMargin              float64  `json:"face_margin"`
		FaceSize                int      `json:"face_size"`
		IgnoreLargerFaces       bool     `json:"ignore_larger_faces"`
		Strength1               float64  `json:"strength1"`
		ApplyInsideMaskOnly     bool     `json:"apply_inside_mask_only"`
		Strength2               float64  `json:"strength2"`
		PromptForFace           string   `json:"prompt_for_face"`
		ModelInputPromptForFace string   `json:"model_input_prompt_for_face"`
	} `json:"face_editor_d_t_o"`
	UltimateUpscaleDTO struct {
		Enabled         bool    `json:"enabled"`
		TargetSizeType  int     `json:"target_size_type"`
		UpscalerIndex   int     `json:"upscaler_index"`
		RedrawMode      int     `json:"redraw_mode"`
		TileWidth       int     `json:"tile_width"`
		TileHeight      int     `json:"tile_height"`
		MaskBlur        int     `json:"mask_blur"`
		SeamsFixType    int     `json:"seams_fix_type"`
		SeamsFixWidth   int     `json:"seams_fix_width"`
		SeamsFixDenoise float64 `json:"seams_fix_denoise"`
		SeamsFixPadding int     `json:"seams_fix_padding"`
	} `json:"ultimate_upscale_d_t_o"`
	AdetailerDTOS []struct {
		AdModel                  string `json:"ad_model"`
		AdNegativePrompt         string `json:"ad_negative_prompt"`
		AdPrompt                 string `json:"ad_prompt"`
		ModelInputNegativePrompt string `json:"model_input_negative_prompt"`
		ModelInputPrompt         string `json:"model_input_prompt"`
	} `json:"adetailer_d_t_o_s"`
}

func (c *CreateImageProRequest) String() string {
	return fmt.Sprintf("%+v", *c)
}

type CreateImageProResponse struct {
	BaseResponse
	Data CreateImageProData `json:"data"`
}

type CreateImageProData struct {
	Results []CreateImageProResult `json:"results"`
}

type CreateImageProResult struct {
	Key            string `json:"key"`
	ExpectedSecond int    `json:"expected_second"`
}

type GeneratingInfoProResponse struct {
	BaseResponse
	Data GeneratingInfoProData `json:"data"`
}

type GeneratingInfoProData struct {
	Infos []GeneratingInfoPro `json:"infos"`
}

type GeneratingInfoPro struct {
	Key             string  `json:"key"`
	Status          int     `json:"status"`
	PictureURL      string  `json:"picture_url"`
	ExpectedSeconds int     `json:"expected_seconds"`
	StartGenTime    int     `json:"start_gen_time"`
	CompleteTime    int     `json:"complete_time"`
	CompletePercent float64 `json:"complete_percent"`
	InvolveYellow   int     `json:"involve_yellow"`
	AuditInfo       string  `json:"audit_info"`
	FailMessage     struct {
		FailCode    int    `json:"fail_code"`
		FailMessage string `json:"fail_message"`
	} `json:"fail_message"`
}

type CreateAvatarArtworkRequest struct {
	AvatarKey        string            `json:"avatar_key"`
	Prompt           string            `json:"prompt"`
	ArtworkTemplates []ArtworkTemplate `json:"artwork_templates"`
	NotifyURL        string            `json:"notify_url"`
}

func (c *CreateAvatarArtworkRequest) String() string {
	return fmt.Sprintf("%+v", *c)
}

type ArtworkTemplate struct {
	TemplateKey string `json:"template_key"`
	Number      int    `json:"number"`
}
type CreateAvatarArtworkResponse struct {
	BaseResponse
	Data CreateAvatarArtworkData `json:"data"`
}

type CreateAvatarArtworkData struct {
	Keys    []string `json:"keys"`
	Results []struct {
		Key            string `json:"key"`
		ExpectedSecond int    `json:"expected_second"`
	} `json:"results"`
	ExpectedIntegralCost int `json:"expected_integral_cost"`
}

type AvatarDefaultResourceResponse struct {
	BaseResponse
	Data AvatarDefaultResource `json:"data"`
}

type AvatarDefaultResource struct {
	AvatarKey       string `json:"avatar_key"`
	TemplateOptions []struct {
		TemplateKey  string `json:"template_key"`
		TemplateName string `json:"template_name"`
		ThemeKey     string `json:"theme_key"`
		ThemeName    string `json:"theme_name"`
	} `json:"template_options"`
}

type CreateSpellAnalysisRequest struct {
	ImageURL  string `json:"image_url"`
	NotifyURL string `json:"notify_url"`
}

func (c *CreateSpellAnalysisRequest) String() string {
	return fmt.Sprintf("%+v", *c)
}

type CreateSpellAnalysisResponse struct {
	BaseResponse
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}

type SpellAnalysisInfoResponse struct {
	BaseResponse
	Data SpellAnalysisInfo `json:"data"`
}

type SpellAnalysisInfo struct {
	SpellAnalysisInfoKey string `json:"spell_analysis_info_key"`
	Tags                 string `json:"tags"`
	ImageURL             string `json:"image_url"`
	Status               int    `json:"status"`
}

type MagicDiceThemeResponse struct {
	BaseResponse
	Data []MagicDiceTheme `json:"data"`
}

type MagicDiceTheme struct {
	ThemeID int    `json:"theme_id"`
	Name    string `json:"name"`
}

type CreateMagicDiceRequest struct {
	Type     CreateMagicDiceType     `json:"type"`
	Model    CreateMagicDiceModel    `json:"model"`
	Keyword  string                  `json:"keyword"`
	ThemeId  int                     `json:"theme_id"`
	Language CreateMagicDiceLanguage `json:"language"`
}

func (c *CreateMagicDiceRequest) String() string {
	return fmt.Sprintf("%+v", *c)
}

type CreateMagicDiceResponse struct {
	BaseResponse
	Data CreateMagicDiceResult `json:"data"`
}

type CreateMagicDiceResult struct {
	PromptChinese string   `json:"prompt_chinese"`
	PromptEnglish string   `json:"prompt_english"`
	Model         string   `json:"model"`
	ModelCode     int      `json:"model_code"`
	Cfg           int      `json:"cfg"`
	ImageType     []string `json:"image_type"`
	Style         []string `json:"style"`
	Artists       []string `json:"artists"`
	ElementMagic  []string `json:"element_magic"`
	Character     []string `json:"character"`
	ModelFusion   []struct {
		Key    string  `json:"key"`
		Weight float64 `json:"weight"`
	} `json:"model_fusion"`
}

type CreateAvatarRequest struct {
	TrainImageUrlList []string `json:"train_image_url_list"`
	NotifyUrl         string   `json:"notify_url"`
}

func (c *CreateAvatarRequest) String() string {
	return fmt.Sprintf("%+v", *c)
}

type CreateAvatarResponse struct {
	BaseResponse
	Data CreateAvatarData `json:"data"`
}

type CreateAvatarData struct {
	Key            string `json:"key"`
	ExpectedSecond int    `json:"expected_second"`
}

type AvatarInfoResponse struct {
	BaseResponse
	Data AvatarInfoData `json:"data"`
}

type AvatarInfoData struct {
	Key             string `json:"key"`
	ModelFusionName string `json:"model_fusion_name"`
	Status          int    `json:"status"`
}

type ImageBatchCheckResponse struct {
	BaseResponse
	Data struct {
		ImageCheckInfoList []ImageCheckInfo `json:"image_check_info_list"`
	} `json:"data"`
}

type ImageCheckInfo struct {
	ImageUrl   string  `json:"image_url"`
	Pass       bool    `json:"pass"`
	Status     string  `json:"status"`
	Message    string  `json:"message"`
	Similarity float64 `json:"similarity"`
}

type CreateVideoRequest struct {
	OriginVideoUrl string `json:"origin_video_url"`
	VideoDuration  int    `json:"video_duration"`
	ModelCode      int    `json:"model_code"`
	QueueType      int    `json:"queue_type"`
	NotifyUrl      string `json:"notify_url"`
}

func (c *CreateVideoRequest) String() string {
	return fmt.Sprintf("%+v", *c)
}

type CreateVideoResponse struct {
	BaseResponse
	Data struct {
		Key string `json:"key"`
	} `json:"data"`
}
