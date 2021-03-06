package structs

import (
	"encoding/json"
	"fmt"
	"time"
)

type AtmosphericTemperature struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}

type HWS struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}

type PRE struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}

type Compass struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}

type Sol struct {
	AT       AtmosphericTemperature `json:"AT"`
	FirstUTC time.Time              `json:"First_UTC"`
	HWS      HWS                    `json:"HWS"`
	LastUTC  time.Time              `json:"Last_UTC"`
	PRE      PRE                    `json:"PRE"`
	Season   string                 `json:"Season"`
	WD       map[string]Compass     `json:"WD"`
}
type FriendlyResponse struct {
	Sols        map[string]Sol `json:"sols"`
	HottestSol  Sol            `json:"hottest_sol"`
	WindiestSol Sol            `json:"windiest_sol"`
	HeaviestSol Sol            `json:"heaviest_sol"`
}

func (swr SolWeatherResponse) ToFriendlyResponse() FriendlyResponse {
	j, err := json.Marshal(swr)
	if err != nil {
		fmt.Printf("error serializing SolWeatherResponse %+v", err)
		return FriendlyResponse{}
	}

	var solsInfo map[string]Sol
	err = json.Unmarshal(j, &solsInfo)
	if err != nil {
		fmt.Printf("error Deserializing SolWeatherResponse Data %+v", err)
		return FriendlyResponse{}
	}

	h, w, heaviest := getTheMostSols(solsInfo)

	return FriendlyResponse{
		Sols:        solsInfo,
		HottestSol:  h,
		WindiestSol: w,
		HeaviestSol: heaviest,
	}
}
func getTheMostSols(solsInfo map[string]Sol) (hottest Sol, windiest Sol, heaviest Sol) {
	for _, v := range solsInfo {
		if hottest.AT.Av == 0 {
			hottest = v
		} else if v.AT.Av > hottest.AT.Av {
			hottest = v
		}
		if v.HWS.Av > windiest.HWS.Av {
			windiest = v
		}
		if v.PRE.Av > heaviest.PRE.Av {
			heaviest = v
		}
	}
	return hottest, windiest, heaviest
}

type SolWeatherResponse struct {
	NUM001 *Sol `json:"1,omitempty"`
	NUM002 *Sol `json:"2,omitempty"`
	NUM003 *Sol `json:"3,omitempty"`
	NUM004 *Sol `json:"4,omitempty"`
	NUM005 *Sol `json:"5,omitempty"`
	NUM006 *Sol `json:"6,omitempty"`
	NUM007 *Sol `json:"7,omitempty"`
	NUM008 *Sol `json:"8,omitempty"`
	NUM009 *Sol `json:"9,omitempty"`
	NUM010 *Sol `json:"10,omitempty"`
	NUM011 *Sol `json:"11,omitempty"`
	NUM012 *Sol `json:"12,omitempty"`
	NUM013 *Sol `json:"13,omitempty"`
	NUM014 *Sol `json:"14,omitempty"`
	NUM015 *Sol `json:"15,omitempty"`
	NUM016 *Sol `json:"16,omitempty"`
	NUM017 *Sol `json:"17,omitempty"`
	NUM018 *Sol `json:"18,omitempty"`
	NUM019 *Sol `json:"19,omitempty"`
	NUM020 *Sol `json:"20,omitempty"`
	NUM021 *Sol `json:"21,omitempty"`
	NUM022 *Sol `json:"22,omitempty"`
	NUM023 *Sol `json:"23,omitempty"`
	NUM024 *Sol `json:"24,omitempty"`
	NUM025 *Sol `json:"25,omitempty"`
	NUM026 *Sol `json:"26,omitempty"`
	NUM027 *Sol `json:"27,omitempty"`
	NUM028 *Sol `json:"28,omitempty"`
	NUM029 *Sol `json:"29,omitempty"`
	NUM030 *Sol `json:"30,omitempty"`
	NUM031 *Sol `json:"31,omitempty"`
	NUM032 *Sol `json:"32,omitempty"`
	NUM033 *Sol `json:"33,omitempty"`
	NUM034 *Sol `json:"34,omitempty"`
	NUM035 *Sol `json:"35,omitempty"`
	NUM036 *Sol `json:"36,omitempty"`
	NUM037 *Sol `json:"37,omitempty"`
	NUM038 *Sol `json:"38,omitempty"`
	NUM039 *Sol `json:"39,omitempty"`
	NUM040 *Sol `json:"40,omitempty"`
	NUM041 *Sol `json:"41,omitempty"`
	NUM042 *Sol `json:"42,omitempty"`
	NUM043 *Sol `json:"43,omitempty"`
	NUM044 *Sol `json:"44,omitempty"`
	NUM045 *Sol `json:"45,omitempty"`
	NUM046 *Sol `json:"46,omitempty"`
	NUM047 *Sol `json:"47,omitempty"`
	NUM048 *Sol `json:"48,omitempty"`
	NUM049 *Sol `json:"49,omitempty"`
	NUM050 *Sol `json:"50,omitempty"`
	NUM051 *Sol `json:"51,omitempty"`
	NUM052 *Sol `json:"52,omitempty"`
	NUM053 *Sol `json:"53,omitempty"`
	NUM054 *Sol `json:"54,omitempty"`
	NUM055 *Sol `json:"55,omitempty"`
	NUM056 *Sol `json:"56,omitempty"`
	NUM057 *Sol `json:"57,omitempty"`
	NUM058 *Sol `json:"58,omitempty"`
	NUM059 *Sol `json:"59,omitempty"`
	NUM060 *Sol `json:"60,omitempty"`
	NUM061 *Sol `json:"61,omitempty"`
	NUM062 *Sol `json:"62,omitempty"`
	NUM063 *Sol `json:"63,omitempty"`
	NUM064 *Sol `json:"64,omitempty"`
	NUM065 *Sol `json:"65,omitempty"`
	NUM066 *Sol `json:"66,omitempty"`
	NUM067 *Sol `json:"67,omitempty"`
	NUM068 *Sol `json:"68,omitempty"`
	NUM069 *Sol `json:"69,omitempty"`
	NUM070 *Sol `json:"70,omitempty"`
	NUM071 *Sol `json:"71,omitempty"`
	NUM072 *Sol `json:"72,omitempty"`
	NUM073 *Sol `json:"73,omitempty"`
	NUM074 *Sol `json:"74,omitempty"`
	NUM075 *Sol `json:"75,omitempty"`
	NUM076 *Sol `json:"76,omitempty"`
	NUM077 *Sol `json:"77,omitempty"`
	NUM078 *Sol `json:"78,omitempty"`
	NUM079 *Sol `json:"79,omitempty"`
	NUM080 *Sol `json:"80,omitempty"`
	NUM081 *Sol `json:"81,omitempty"`
	NUM082 *Sol `json:"82,omitempty"`
	NUM083 *Sol `json:"83,omitempty"`
	NUM084 *Sol `json:"84,omitempty"`
	NUM085 *Sol `json:"85,omitempty"`
	NUM086 *Sol `json:"86,omitempty"`
	NUM087 *Sol `json:"87,omitempty"`
	NUM088 *Sol `json:"88,omitempty"`
	NUM089 *Sol `json:"89,omitempty"`
	NUM090 *Sol `json:"90,omitempty"`
	NUM091 *Sol `json:"91,omitempty"`
	NUM092 *Sol `json:"92,omitempty"`
	NUM093 *Sol `json:"93,omitempty"`
	NUM094 *Sol `json:"94,omitempty"`
	NUM095 *Sol `json:"95,omitempty"`
	NUM096 *Sol `json:"96,omitempty"`
	NUM097 *Sol `json:"97,omitempty"`
	NUM098 *Sol `json:"98,omitempty"`
	NUM099 *Sol `json:"99,omitempty"`
	NUM100 *Sol `json:"100,omitempty"`
	NUM101 *Sol `json:"101,omitempty"`
	NUM102 *Sol `json:"102,omitempty"`
	NUM103 *Sol `json:"103,omitempty"`
	NUM104 *Sol `json:"104,omitempty"`
	NUM105 *Sol `json:"105,omitempty"`
	NUM106 *Sol `json:"106,omitempty"`
	NUM107 *Sol `json:"107,omitempty"`
	NUM108 *Sol `json:"108,omitempty"`
	NUM109 *Sol `json:"109,omitempty"`
	NUM110 *Sol `json:"110,omitempty"`
	NUM111 *Sol `json:"111,omitempty"`
	NUM112 *Sol `json:"112,omitempty"`
	NUM113 *Sol `json:"113,omitempty"`
	NUM114 *Sol `json:"114,omitempty"`
	NUM115 *Sol `json:"115,omitempty"`
	NUM116 *Sol `json:"116,omitempty"`
	NUM117 *Sol `json:"117,omitempty"`
	NUM118 *Sol `json:"118,omitempty"`
	NUM119 *Sol `json:"119,omitempty"`
	NUM120 *Sol `json:"120,omitempty"`
	NUM121 *Sol `json:"121,omitempty"`
	NUM122 *Sol `json:"122,omitempty"`
	NUM123 *Sol `json:"123,omitempty"`
	NUM124 *Sol `json:"124,omitempty"`
	NUM125 *Sol `json:"125,omitempty"`
	NUM126 *Sol `json:"126,omitempty"`
	NUM127 *Sol `json:"127,omitempty"`
	NUM128 *Sol `json:"128,omitempty"`
	NUM129 *Sol `json:"129,omitempty"`
	NUM130 *Sol `json:"130,omitempty"`
	NUM131 *Sol `json:"131,omitempty"`
	NUM132 *Sol `json:"132,omitempty"`
	NUM133 *Sol `json:"133,omitempty"`
	NUM134 *Sol `json:"134,omitempty"`
	NUM135 *Sol `json:"135,omitempty"`
	NUM136 *Sol `json:"136,omitempty"`
	NUM137 *Sol `json:"137,omitempty"`
	NUM138 *Sol `json:"138,omitempty"`
	NUM139 *Sol `json:"139,omitempty"`
	NUM140 *Sol `json:"140,omitempty"`
	NUM141 *Sol `json:"141,omitempty"`
	NUM142 *Sol `json:"142,omitempty"`
	NUM143 *Sol `json:"143,omitempty"`
	NUM144 *Sol `json:"144,omitempty"`
	NUM145 *Sol `json:"145,omitempty"`
	NUM146 *Sol `json:"146,omitempty"`
	NUM147 *Sol `json:"147,omitempty"`
	NUM148 *Sol `json:"148,omitempty"`
	NUM149 *Sol `json:"149,omitempty"`
	NUM150 *Sol `json:"150,omitempty"`
	NUM151 *Sol `json:"151,omitempty"`
	NUM152 *Sol `json:"152,omitempty"`
	NUM153 *Sol `json:"153,omitempty"`
	NUM154 *Sol `json:"154,omitempty"`
	NUM155 *Sol `json:"155,omitempty"`
	NUM156 *Sol `json:"156,omitempty"`
	NUM157 *Sol `json:"157,omitempty"`
	NUM158 *Sol `json:"158,omitempty"`
	NUM159 *Sol `json:"159,omitempty"`
	NUM160 *Sol `json:"160,omitempty"`
	NUM161 *Sol `json:"161,omitempty"`
	NUM162 *Sol `json:"162,omitempty"`
	NUM163 *Sol `json:"163,omitempty"`
	NUM164 *Sol `json:"164,omitempty"`
	NUM165 *Sol `json:"165,omitempty"`
	NUM166 *Sol `json:"166,omitempty"`
	NUM167 *Sol `json:"167,omitempty"`
	NUM168 *Sol `json:"168,omitempty"`
	NUM169 *Sol `json:"169,omitempty"`
	NUM170 *Sol `json:"170,omitempty"`
	NUM171 *Sol `json:"171,omitempty"`
	NUM172 *Sol `json:"172,omitempty"`
	NUM173 *Sol `json:"173,omitempty"`
	NUM174 *Sol `json:"174,omitempty"`
	NUM175 *Sol `json:"175,omitempty"`
	NUM176 *Sol `json:"176,omitempty"`
	NUM177 *Sol `json:"177,omitempty"`
	NUM178 *Sol `json:"178,omitempty"`
	NUM179 *Sol `json:"179,omitempty"`
	NUM180 *Sol `json:"180,omitempty"`
	NUM181 *Sol `json:"181,omitempty"`
	NUM182 *Sol `json:"182,omitempty"`
	NUM183 *Sol `json:"183,omitempty"`
	NUM184 *Sol `json:"184,omitempty"`
	NUM185 *Sol `json:"185,omitempty"`
	NUM186 *Sol `json:"186,omitempty"`
	NUM187 *Sol `json:"187,omitempty"`
	NUM188 *Sol `json:"188,omitempty"`
	NUM189 *Sol `json:"189,omitempty"`
	NUM190 *Sol `json:"190,omitempty"`
	NUM191 *Sol `json:"191,omitempty"`
	NUM192 *Sol `json:"192,omitempty"`
	NUM193 *Sol `json:"193,omitempty"`
	NUM194 *Sol `json:"194,omitempty"`
	NUM195 *Sol `json:"195,omitempty"`
	NUM196 *Sol `json:"196,omitempty"`
	NUM197 *Sol `json:"197,omitempty"`
	NUM198 *Sol `json:"198,omitempty"`
	NUM199 *Sol `json:"199,omitempty"`
	NUM200 *Sol `json:"200,omitempty"`
	NUM201 *Sol `json:"201,omitempty"`
	NUM202 *Sol `json:"202,omitempty"`
	NUM203 *Sol `json:"203,omitempty"`
	NUM204 *Sol `json:"204,omitempty"`
	NUM205 *Sol `json:"205,omitempty"`
	NUM206 *Sol `json:"206,omitempty"`
	NUM207 *Sol `json:"207,omitempty"`
	NUM208 *Sol `json:"208,omitempty"`
	NUM209 *Sol `json:"209,omitempty"`
	NUM210 *Sol `json:"210,omitempty"`
	NUM211 *Sol `json:"211,omitempty"`
	NUM212 *Sol `json:"212,omitempty"`
	NUM213 *Sol `json:"213,omitempty"`
	NUM214 *Sol `json:"214,omitempty"`
	NUM215 *Sol `json:"215,omitempty"`
	NUM216 *Sol `json:"216,omitempty"`
	NUM217 *Sol `json:"217,omitempty"`
	NUM218 *Sol `json:"218,omitempty"`
	NUM219 *Sol `json:"219,omitempty"`
	NUM220 *Sol `json:"220,omitempty"`
	NUM221 *Sol `json:"221,omitempty"`
	NUM222 *Sol `json:"222,omitempty"`
	NUM223 *Sol `json:"223,omitempty"`
	NUM224 *Sol `json:"224,omitempty"`
	NUM225 *Sol `json:"225,omitempty"`
	NUM226 *Sol `json:"226,omitempty"`
	NUM227 *Sol `json:"227,omitempty"`
	NUM228 *Sol `json:"228,omitempty"`
	NUM229 *Sol `json:"229,omitempty"`
	NUM230 *Sol `json:"230,omitempty"`
	NUM231 *Sol `json:"231,omitempty"`
	NUM232 *Sol `json:"232,omitempty"`
	NUM233 *Sol `json:"233,omitempty"`
	NUM234 *Sol `json:"234,omitempty"`
	NUM235 *Sol `json:"235,omitempty"`
	NUM236 *Sol `json:"236,omitempty"`
	NUM237 *Sol `json:"237,omitempty"`
	NUM238 *Sol `json:"238,omitempty"`
	NUM239 *Sol `json:"239,omitempty"`
	NUM240 *Sol `json:"240,omitempty"`
	NUM241 *Sol `json:"241,omitempty"`
	NUM242 *Sol `json:"242,omitempty"`
	NUM243 *Sol `json:"243,omitempty"`
	NUM244 *Sol `json:"244,omitempty"`
	NUM245 *Sol `json:"245,omitempty"`
	NUM246 *Sol `json:"246,omitempty"`
	NUM247 *Sol `json:"247,omitempty"`
	NUM248 *Sol `json:"248,omitempty"`
	NUM249 *Sol `json:"249,omitempty"`
	NUM250 *Sol `json:"250,omitempty"`
	NUM251 *Sol `json:"251,omitempty"`
	NUM252 *Sol `json:"252,omitempty"`
	NUM253 *Sol `json:"253,omitempty"`
	NUM254 *Sol `json:"254,omitempty"`
	NUM255 *Sol `json:"255,omitempty"`
	NUM256 *Sol `json:"256,omitempty"`
	NUM257 *Sol `json:"257,omitempty"`
	NUM258 *Sol `json:"258,omitempty"`
	NUM259 *Sol `json:"259,omitempty"`
	NUM260 *Sol `json:"260,omitempty"`
	NUM261 *Sol `json:"261,omitempty"`
	NUM262 *Sol `json:"262,omitempty"`
	NUM263 *Sol `json:"263,omitempty"`
	NUM264 *Sol `json:"264,omitempty"`
	NUM265 *Sol `json:"265,omitempty"`
	NUM266 *Sol `json:"266,omitempty"`
	NUM267 *Sol `json:"267,omitempty"`
	NUM268 *Sol `json:"268,omitempty"`
	NUM269 *Sol `json:"269,omitempty"`
	NUM270 *Sol `json:"270,omitempty"`
	NUM271 *Sol `json:"271,omitempty"`
	NUM272 *Sol `json:"272,omitempty"`
	NUM273 *Sol `json:"273,omitempty"`
	NUM274 *Sol `json:"274,omitempty"`
	NUM275 *Sol `json:"275,omitempty"`
	NUM276 *Sol `json:"276,omitempty"`
	NUM277 *Sol `json:"277,omitempty"`
	NUM278 *Sol `json:"278,omitempty"`
	NUM279 *Sol `json:"279,omitempty"`
	NUM280 *Sol `json:"280,omitempty"`
	NUM281 *Sol `json:"281,omitempty"`
	NUM282 *Sol `json:"282,omitempty"`
	NUM283 *Sol `json:"283,omitempty"`
	NUM284 *Sol `json:"284,omitempty"`
	NUM285 *Sol `json:"285,omitempty"`
	NUM286 *Sol `json:"286,omitempty"`
	NUM287 *Sol `json:"287,omitempty"`
	NUM288 *Sol `json:"288,omitempty"`
	NUM289 *Sol `json:"289,omitempty"`
	NUM290 *Sol `json:"290,omitempty"`
	NUM291 *Sol `json:"291,omitempty"`
	NUM292 *Sol `json:"292,omitempty"`
	NUM293 *Sol `json:"293,omitempty"`
	NUM294 *Sol `json:"294,omitempty"`
	NUM295 *Sol `json:"295,omitempty"`
	NUM296 *Sol `json:"296,omitempty"`
	NUM297 *Sol `json:"297,omitempty"`
	NUM298 *Sol `json:"298,omitempty"`
	NUM299 *Sol `json:"299,omitempty"`
	NUM300 *Sol `json:"300,omitempty"`
	NUM301 *Sol `json:"301,omitempty"`
	NUM302 *Sol `json:"302,omitempty"`
	NUM303 *Sol `json:"303,omitempty"`
	NUM304 *Sol `json:"304,omitempty"`
	NUM305 *Sol `json:"305,omitempty"`
	NUM306 *Sol `json:"306,omitempty"`
	NUM307 *Sol `json:"307,omitempty"`
	NUM308 *Sol `json:"308,omitempty"`
	NUM309 *Sol `json:"309,omitempty"`
	NUM310 *Sol `json:"310,omitempty"`
	NUM311 *Sol `json:"311,omitempty"`
	NUM312 *Sol `json:"312,omitempty"`
	NUM313 *Sol `json:"313,omitempty"`
	NUM314 *Sol `json:"314,omitempty"`
	NUM315 *Sol `json:"315,omitempty"`
	NUM316 *Sol `json:"316,omitempty"`
	NUM317 *Sol `json:"317,omitempty"`
	NUM318 *Sol `json:"318,omitempty"`
	NUM319 *Sol `json:"319,omitempty"`
	NUM320 *Sol `json:"320,omitempty"`
	NUM321 *Sol `json:"321,omitempty"`
	NUM322 *Sol `json:"322,omitempty"`
	NUM323 *Sol `json:"323,omitempty"`
	NUM324 *Sol `json:"324,omitempty"`
	NUM325 *Sol `json:"325,omitempty"`
	NUM326 *Sol `json:"326,omitempty"`
	NUM327 *Sol `json:"327,omitempty"`
	NUM328 *Sol `json:"328,omitempty"`
	NUM329 *Sol `json:"329,omitempty"`
	NUM330 *Sol `json:"330,omitempty"`
	NUM331 *Sol `json:"331,omitempty"`
	NUM332 *Sol `json:"332,omitempty"`
	NUM333 *Sol `json:"333,omitempty"`
	NUM334 *Sol `json:"334,omitempty"`
	NUM335 *Sol `json:"335,omitempty"`
	NUM336 *Sol `json:"336,omitempty"`
	NUM337 *Sol `json:"337,omitempty"`
	NUM338 *Sol `json:"338,omitempty"`
	NUM339 *Sol `json:"339,omitempty"`
	NUM340 *Sol `json:"340,omitempty"`
	NUM341 *Sol `json:"341,omitempty"`
	NUM342 *Sol `json:"342,omitempty"`
	NUM343 *Sol `json:"343,omitempty"`
	NUM344 *Sol `json:"344,omitempty"`
	NUM345 *Sol `json:"345,omitempty"`
	NUM346 *Sol `json:"346,omitempty"`
	NUM347 *Sol `json:"347,omitempty"`
	NUM348 *Sol `json:"348,omitempty"`
	NUM349 *Sol `json:"349,omitempty"`
	NUM350 *Sol `json:"350,omitempty"`
	NUM351 *Sol `json:"351,omitempty"`
	NUM352 *Sol `json:"352,omitempty"`
	NUM353 *Sol `json:"353,omitempty"`
	NUM354 *Sol `json:"354,omitempty"`
	NUM355 *Sol `json:"355,omitempty"`
	NUM356 *Sol `json:"356,omitempty"`
	NUM357 *Sol `json:"357,omitempty"`
	NUM358 *Sol `json:"358,omitempty"`
	NUM359 *Sol `json:"359,omitempty"`
	NUM360 *Sol `json:"360,omitempty"`
	NUM361 *Sol `json:"361,omitempty"`
	NUM362 *Sol `json:"362,omitempty"`
	NUM363 *Sol `json:"363,omitempty"`
	NUM364 *Sol `json:"364,omitempty"`
	NUM365 *Sol `json:"365,omitempty"`
	NUM366 *Sol `json:"366,omitempty"`
	NUM367 *Sol `json:"367,omitempty"`
	NUM368 *Sol `json:"368,omitempty"`
	NUM369 *Sol `json:"369,omitempty"`
	NUM370 *Sol `json:"370,omitempty"`
	NUM371 *Sol `json:"371,omitempty"`
	NUM372 *Sol `json:"372,omitempty"`
	NUM373 *Sol `json:"373,omitempty"`
	NUM374 *Sol `json:"374,omitempty"`
	NUM375 *Sol `json:"375,omitempty"`
	NUM376 *Sol `json:"376,omitempty"`
	NUM377 *Sol `json:"377,omitempty"`
	NUM378 *Sol `json:"378,omitempty"`
	NUM379 *Sol `json:"379,omitempty"`
	NUM380 *Sol `json:"380,omitempty"`
	NUM381 *Sol `json:"381,omitempty"`
	NUM382 *Sol `json:"382,omitempty"`
	NUM383 *Sol `json:"383,omitempty"`
	NUM384 *Sol `json:"384,omitempty"`
	NUM385 *Sol `json:"385,omitempty"`
	NUM386 *Sol `json:"386,omitempty"`
	NUM387 *Sol `json:"387,omitempty"`
	NUM388 *Sol `json:"388,omitempty"`
	NUM389 *Sol `json:"389,omitempty"`
	NUM390 *Sol `json:"390,omitempty"`
	NUM391 *Sol `json:"391,omitempty"`
	NUM392 *Sol `json:"392,omitempty"`
	NUM393 *Sol `json:"393,omitempty"`
	NUM394 *Sol `json:"394,omitempty"`
	NUM395 *Sol `json:"395,omitempty"`
	NUM396 *Sol `json:"396,omitempty"`
	NUM397 *Sol `json:"397,omitempty"`
	NUM398 *Sol `json:"398,omitempty"`
	NUM399 *Sol `json:"399,omitempty"`
	NUM400 *Sol `json:"400,omitempty"`
	NUM401 *Sol `json:"401,omitempty"`
	NUM402 *Sol `json:"402,omitempty"`
	NUM403 *Sol `json:"403,omitempty"`
	NUM404 *Sol `json:"404,omitempty"`
	NUM405 *Sol `json:"405,omitempty"`
	NUM406 *Sol `json:"406,omitempty"`
	NUM407 *Sol `json:"407,omitempty"`
	NUM408 *Sol `json:"408,omitempty"`
	NUM409 *Sol `json:"409,omitempty"`
	NUM410 *Sol `json:"410,omitempty"`
	NUM411 *Sol `json:"411,omitempty"`
	NUM412 *Sol `json:"412,omitempty"`
	NUM413 *Sol `json:"413,omitempty"`
	NUM414 *Sol `json:"414,omitempty"`
	NUM415 *Sol `json:"415,omitempty"`
	NUM416 *Sol `json:"416,omitempty"`
	NUM417 *Sol `json:"417,omitempty"`
	NUM418 *Sol `json:"418,omitempty"`
	NUM419 *Sol `json:"419,omitempty"`
	NUM420 *Sol `json:"420,omitempty"`
	NUM421 *Sol `json:"421,omitempty"`
	NUM422 *Sol `json:"422,omitempty"`
	NUM423 *Sol `json:"423,omitempty"`
	NUM424 *Sol `json:"424,omitempty"`
	NUM425 *Sol `json:"425,omitempty"`
	NUM426 *Sol `json:"426,omitempty"`
	NUM427 *Sol `json:"427,omitempty"`
	NUM428 *Sol `json:"428,omitempty"`
	NUM429 *Sol `json:"429,omitempty"`
	NUM430 *Sol `json:"430,omitempty"`
	NUM431 *Sol `json:"431,omitempty"`
	NUM432 *Sol `json:"432,omitempty"`
	NUM433 *Sol `json:"433,omitempty"`
	NUM434 *Sol `json:"434,omitempty"`
	NUM435 *Sol `json:"435,omitempty"`
	NUM436 *Sol `json:"436,omitempty"`
	NUM437 *Sol `json:"437,omitempty"`
	NUM438 *Sol `json:"438,omitempty"`
	NUM439 *Sol `json:"439,omitempty"`
	NUM440 *Sol `json:"440,omitempty"`
	NUM441 *Sol `json:"441,omitempty"`
	NUM442 *Sol `json:"442,omitempty"`
	NUM443 *Sol `json:"443,omitempty"`
	NUM444 *Sol `json:"444,omitempty"`
	NUM445 *Sol `json:"445,omitempty"`
	NUM446 *Sol `json:"446,omitempty"`
	NUM447 *Sol `json:"447,omitempty"`
	NUM448 *Sol `json:"448,omitempty"`
	NUM449 *Sol `json:"449,omitempty"`
	NUM450 *Sol `json:"450,omitempty"`
	NUM451 *Sol `json:"451,omitempty"`
	NUM452 *Sol `json:"452,omitempty"`
	NUM453 *Sol `json:"453,omitempty"`
	NUM454 *Sol `json:"454,omitempty"`
	NUM455 *Sol `json:"455,omitempty"`
	NUM456 *Sol `json:"456,omitempty"`
	NUM457 *Sol `json:"457,omitempty"`
	NUM458 *Sol `json:"458,omitempty"`
	NUM459 *Sol `json:"459,omitempty"`
	NUM460 *Sol `json:"460,omitempty"`
	NUM461 *Sol `json:"461,omitempty"`
	NUM462 *Sol `json:"462,omitempty"`
	NUM463 *Sol `json:"463,omitempty"`
	NUM464 *Sol `json:"464,omitempty"`
	NUM465 *Sol `json:"465,omitempty"`
	NUM466 *Sol `json:"466,omitempty"`
	NUM467 *Sol `json:"467,omitempty"`
	NUM468 *Sol `json:"468,omitempty"`
	NUM469 *Sol `json:"469,omitempty"`
	NUM470 *Sol `json:"470,omitempty"`
	NUM471 *Sol `json:"471,omitempty"`
	NUM472 *Sol `json:"472,omitempty"`
	NUM473 *Sol `json:"473,omitempty"`
	NUM474 *Sol `json:"474,omitempty"`
	NUM475 *Sol `json:"475,omitempty"`
	NUM476 *Sol `json:"476,omitempty"`
	NUM477 *Sol `json:"477,omitempty"`
	NUM478 *Sol `json:"478,omitempty"`
	NUM479 *Sol `json:"479,omitempty"`
	NUM480 *Sol `json:"480,omitempty"`
	NUM481 *Sol `json:"481,omitempty"`
	NUM482 *Sol `json:"482,omitempty"`
	NUM483 *Sol `json:"483,omitempty"`
	NUM484 *Sol `json:"484,omitempty"`
	NUM485 *Sol `json:"485,omitempty"`
	NUM486 *Sol `json:"486,omitempty"`
	NUM487 *Sol `json:"487,omitempty"`
	NUM488 *Sol `json:"488,omitempty"`
	NUM489 *Sol `json:"489,omitempty"`
	NUM490 *Sol `json:"490,omitempty"`
	NUM491 *Sol `json:"491,omitempty"`
	NUM492 *Sol `json:"492,omitempty"`
	NUM493 *Sol `json:"493,omitempty"`
	NUM494 *Sol `json:"494,omitempty"`
	NUM495 *Sol `json:"495,omitempty"`
	NUM496 *Sol `json:"496,omitempty"`
	NUM497 *Sol `json:"497,omitempty"`
	NUM498 *Sol `json:"498,omitempty"`
	NUM499 *Sol `json:"499,omitempty"`
	NUM500 *Sol `json:"500,omitempty"`
	NUM501 *Sol `json:"501,omitempty"`
	NUM502 *Sol `json:"502,omitempty"`
	NUM503 *Sol `json:"503,omitempty"`
	NUM504 *Sol `json:"504,omitempty"`
	NUM505 *Sol `json:"505,omitempty"`
	NUM506 *Sol `json:"506,omitempty"`
	NUM507 *Sol `json:"507,omitempty"`
	NUM508 *Sol `json:"508,omitempty"`
	NUM509 *Sol `json:"509,omitempty"`
	NUM510 *Sol `json:"510,omitempty"`
	NUM511 *Sol `json:"511,omitempty"`
	NUM512 *Sol `json:"512,omitempty"`
	NUM513 *Sol `json:"513,omitempty"`
	NUM514 *Sol `json:"514,omitempty"`
	NUM515 *Sol `json:"515,omitempty"`
	NUM516 *Sol `json:"516,omitempty"`
	NUM517 *Sol `json:"517,omitempty"`
	NUM518 *Sol `json:"518,omitempty"`
	NUM519 *Sol `json:"519,omitempty"`
	NUM520 *Sol `json:"520,omitempty"`
	NUM521 *Sol `json:"521,omitempty"`
	NUM522 *Sol `json:"522,omitempty"`
	NUM523 *Sol `json:"523,omitempty"`
	NUM524 *Sol `json:"524,omitempty"`
	NUM525 *Sol `json:"525,omitempty"`
	NUM526 *Sol `json:"526,omitempty"`
	NUM527 *Sol `json:"527,omitempty"`
	NUM528 *Sol `json:"528,omitempty"`
	NUM529 *Sol `json:"529,omitempty"`
	NUM530 *Sol `json:"530,omitempty"`
	NUM531 *Sol `json:"531,omitempty"`
	NUM532 *Sol `json:"532,omitempty"`
	NUM533 *Sol `json:"533,omitempty"`
	NUM534 *Sol `json:"534,omitempty"`
	NUM535 *Sol `json:"535,omitempty"`
	NUM536 *Sol `json:"536,omitempty"`
	NUM537 *Sol `json:"537,omitempty"`
	NUM538 *Sol `json:"538,omitempty"`
	NUM539 *Sol `json:"539,omitempty"`
	NUM540 *Sol `json:"540,omitempty"`
	NUM541 *Sol `json:"541,omitempty"`
	NUM542 *Sol `json:"542,omitempty"`
	NUM543 *Sol `json:"543,omitempty"`
	NUM544 *Sol `json:"544,omitempty"`
	NUM545 *Sol `json:"545,omitempty"`
	NUM546 *Sol `json:"546,omitempty"`
	NUM547 *Sol `json:"547,omitempty"`
	NUM548 *Sol `json:"548,omitempty"`
	NUM549 *Sol `json:"549,omitempty"`
	NUM550 *Sol `json:"550,omitempty"`
	NUM551 *Sol `json:"551,omitempty"`
	NUM552 *Sol `json:"552,omitempty"`
	NUM553 *Sol `json:"553,omitempty"`
	NUM554 *Sol `json:"554,omitempty"`
	NUM555 *Sol `json:"555,omitempty"`
	NUM556 *Sol `json:"556,omitempty"`
	NUM557 *Sol `json:"557,omitempty"`
	NUM558 *Sol `json:"558,omitempty"`
	NUM559 *Sol `json:"559,omitempty"`
	NUM560 *Sol `json:"560,omitempty"`
	NUM561 *Sol `json:"561,omitempty"`
	NUM562 *Sol `json:"562,omitempty"`
	NUM563 *Sol `json:"563,omitempty"`
	NUM564 *Sol `json:"564,omitempty"`
	NUM565 *Sol `json:"565,omitempty"`
	NUM566 *Sol `json:"566,omitempty"`
	NUM567 *Sol `json:"567,omitempty"`
	NUM568 *Sol `json:"568,omitempty"`
	NUM569 *Sol `json:"569,omitempty"`
	NUM570 *Sol `json:"570,omitempty"`
	NUM571 *Sol `json:"571,omitempty"`
	NUM572 *Sol `json:"572,omitempty"`
	NUM573 *Sol `json:"573,omitempty"`
	NUM574 *Sol `json:"574,omitempty"`
	NUM575 *Sol `json:"575,omitempty"`
	NUM576 *Sol `json:"576,omitempty"`
	NUM577 *Sol `json:"577,omitempty"`
	NUM578 *Sol `json:"578,omitempty"`
	NUM579 *Sol `json:"579,omitempty"`
	NUM580 *Sol `json:"580,omitempty"`
	NUM581 *Sol `json:"581,omitempty"`
	NUM582 *Sol `json:"582,omitempty"`
	NUM583 *Sol `json:"583,omitempty"`
	NUM584 *Sol `json:"584,omitempty"`
	NUM585 *Sol `json:"585,omitempty"`
	NUM586 *Sol `json:"586,omitempty"`
	NUM587 *Sol `json:"587,omitempty"`
	NUM588 *Sol `json:"588,omitempty"`
	NUM589 *Sol `json:"589,omitempty"`
	NUM590 *Sol `json:"590,omitempty"`
	NUM591 *Sol `json:"591,omitempty"`
	NUM592 *Sol `json:"592,omitempty"`
	NUM593 *Sol `json:"593,omitempty"`
	NUM594 *Sol `json:"594,omitempty"`
	NUM595 *Sol `json:"595,omitempty"`
	NUM596 *Sol `json:"596,omitempty"`
	NUM597 *Sol `json:"597,omitempty"`
	NUM598 *Sol `json:"598,omitempty"`
	NUM599 *Sol `json:"599,omitempty"`
	NUM600 *Sol `json:"600,omitempty"`
	NUM601 *Sol `json:"601,omitempty"`
	NUM602 *Sol `json:"602,omitempty"`
	NUM603 *Sol `json:"603,omitempty"`
	NUM604 *Sol `json:"604,omitempty"`
	NUM605 *Sol `json:"605,omitempty"`
	NUM606 *Sol `json:"606,omitempty"`
	NUM607 *Sol `json:"607,omitempty"`
	NUM608 *Sol `json:"608,omitempty"`
	NUM609 *Sol `json:"609,omitempty"`
	NUM610 *Sol `json:"610,omitempty"`
	NUM611 *Sol `json:"611,omitempty"`
	NUM612 *Sol `json:"612,omitempty"`
	NUM613 *Sol `json:"613,omitempty"`
	NUM614 *Sol `json:"614,omitempty"`
	NUM615 *Sol `json:"615,omitempty"`
	NUM616 *Sol `json:"616,omitempty"`
	NUM617 *Sol `json:"617,omitempty"`
	NUM618 *Sol `json:"618,omitempty"`
	NUM619 *Sol `json:"619,omitempty"`
	NUM620 *Sol `json:"620,omitempty"`
	NUM621 *Sol `json:"621,omitempty"`
	NUM622 *Sol `json:"622,omitempty"`
	NUM623 *Sol `json:"623,omitempty"`
	NUM624 *Sol `json:"624,omitempty"`
	NUM625 *Sol `json:"625,omitempty"`
	NUM626 *Sol `json:"626,omitempty"`
	NUM627 *Sol `json:"627,omitempty"`
	NUM628 *Sol `json:"628,omitempty"`
	NUM629 *Sol `json:"629,omitempty"`
	NUM630 *Sol `json:"630,omitempty"`
	NUM631 *Sol `json:"631,omitempty"`
	NUM632 *Sol `json:"632,omitempty"`
	NUM633 *Sol `json:"633,omitempty"`
	NUM634 *Sol `json:"634,omitempty"`
	NUM635 *Sol `json:"635,omitempty"`
	NUM636 *Sol `json:"636,omitempty"`
	NUM637 *Sol `json:"637,omitempty"`
	NUM638 *Sol `json:"638,omitempty"`
	NUM639 *Sol `json:"639,omitempty"`
	NUM640 *Sol `json:"640,omitempty"`
	NUM641 *Sol `json:"641,omitempty"`
	NUM642 *Sol `json:"642,omitempty"`
	NUM643 *Sol `json:"643,omitempty"`
	NUM644 *Sol `json:"644,omitempty"`
	NUM645 *Sol `json:"645,omitempty"`
	NUM646 *Sol `json:"646,omitempty"`
	NUM647 *Sol `json:"647,omitempty"`
	NUM648 *Sol `json:"648,omitempty"`
	NUM649 *Sol `json:"649,omitempty"`
	NUM650 *Sol `json:"650,omitempty"`
	NUM651 *Sol `json:"651,omitempty"`
	NUM652 *Sol `json:"652,omitempty"`
	NUM653 *Sol `json:"653,omitempty"`
	NUM654 *Sol `json:"654,omitempty"`
	NUM655 *Sol `json:"655,omitempty"`
	NUM656 *Sol `json:"656,omitempty"`
	NUM657 *Sol `json:"657,omitempty"`
	NUM658 *Sol `json:"658,omitempty"`
	NUM659 *Sol `json:"659,omitempty"`
	NUM660 *Sol `json:"660,omitempty"`
	NUM661 *Sol `json:"661,omitempty"`
	NUM662 *Sol `json:"662,omitempty"`
	NUM663 *Sol `json:"663,omitempty"`
	NUM664 *Sol `json:"664,omitempty"`
	NUM665 *Sol `json:"665,omitempty"`
	NUM666 *Sol `json:"666,omitempty"`
	NUM667 *Sol `json:"667,omitempty"`
	NUM668 *Sol `json:"668,omitempty"`
	NUM669 *Sol `json:"669,omitempty"`
	NUM670 *Sol `json:"670,omitempty"`
	NUM671 *Sol `json:"671,omitempty"`
	NUM672 *Sol `json:"672,omitempty"`
	NUM673 *Sol `json:"673,omitempty"`
	NUM674 *Sol `json:"674,omitempty"`
	NUM675 *Sol `json:"675,omitempty"`
	NUM676 *Sol `json:"676,omitempty"`
	NUM677 *Sol `json:"677,omitempty"`
	NUM678 *Sol `json:"678,omitempty"`
	NUM679 *Sol `json:"679,omitempty"`
	NUM680 *Sol `json:"680,omitempty"`
	NUM681 *Sol `json:"681,omitempty"`
	NUM682 *Sol `json:"682,omitempty"`
	NUM683 *Sol `json:"683,omitempty"`
	NUM684 *Sol `json:"684,omitempty"`
	NUM685 *Sol `json:"685,omitempty"`
	NUM686 *Sol `json:"686,omitempty"`
	NUM687 *Sol `json:"687,omitempty"`
}
