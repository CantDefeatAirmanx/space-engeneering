package mocks

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	inventory_v1 "github.com/CantDefeatAirmanx/space-engeneering/shared/pkg/proto/inventory/v1"
)

var InitialParts = []inventory_v1.Part{
	// CATEGORY_ENGINE - Детали для двигателей
	{
		Uuid:          "engine-001",
		Name:          "Турбонасосный агрегат РД-107",
		Description:   "Основной турбонасосный агрегат для ракетного двигателя РД-107. Обеспечивает подачу топлива под высоким давлением.",
		Price:         2500000.0,
		StockQuantity: 5,
		Category:      inventory_v1.Category_CATEGORY_ENGINE,
		Dimensions: &inventory_v1.Dimensions{
			Length: 120.0,
			Width:  80.0,
			Height: 60.0,
			Weight: 450.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "НПО Энергомаш",
			Country: "Россия",
			Website: "https://www.npoenergomash.ru",
		},
		Tags:      []string{"турбонасос", "высокое давление", "критический компонент"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"максимальное_давление": 300.0,
					"обороты_в_минуту":      15000.0,
					"мощность_квт":          500.0,
				},
			},
		},
	},
	{
		Uuid:          "engine-002",
		Name:          "Камера сгорания РД-180",
		Description:   "Камера сгорания для кислородно-керосинового двигателя РД-180. Изготовлена из жаропрочных сплавов.",
		Price:         1800000.0,
		StockQuantity: 8,
		Category:      inventory_v1.Category_CATEGORY_ENGINE,
		Dimensions: &inventory_v1.Dimensions{
			Length: 200.0,
			Width:  150.0,
			Height: 100.0,
			Weight: 800.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "НПО Энергомаш",
			Country: "Россия",
			Website: "https://www.npoenergomash.ru",
		},
		Tags:      []string{"камера сгорания", "жаропрочные сплавы", "высокая температура"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 20, 14, 15, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 20, 14, 15, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"температура_сгорания": 3500.0,
					"давление_в_камере":    250.0,
					"толщина_стенки_мм":    5.0,
				},
			},
		},
	},
	{
		Uuid:          "engine-003",
		Name:          "Сопло Лаваля Merlin-1D",
		Description:   "Сопло Лаваля для двигателя Merlin-1D. Оптимизировано для максимальной тяги в вакууме.",
		Price:         1200000.0,
		StockQuantity: 12,
		Category:      inventory_v1.Category_CATEGORY_ENGINE,
		Dimensions: &inventory_v1.Dimensions{
			Length: 180.0,
			Width:  120.0,
			Height: 90.0,
			Weight: 350.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "SpaceX",
			Country: "США",
			Website: "https://www.spacex.com",
		},
		Tags:      []string{"сопло", "вакуум", "оптимизация тяги"},
		CreatedAt: timestamppb.New(time.Date(2024, 2, 5, 9, 45, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 2, 5, 9, 45, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"коэффициент_расширения": 16.0,
					"угол_раскрытия_градусы": 15.0,
					"длина_сопла_мм":         1800.0,
				},
			},
		},
	},

	// CATEGORY_FUEL - Детали для топлива
	{
		Uuid:          "fuel-001",
		Name:          "Топливный бак RP-1",
		Description:   "Топливный бак для керосина RP-1. Изготовлен из алюминиевого сплава с внутренним покрытием.",
		Price:         800000.0,
		StockQuantity: 15,
		Category:      inventory_v1.Category_CATEGORY_FUEL,
		Dimensions: &inventory_v1.Dimensions{
			Length: 300.0,
			Width:  200.0,
			Height: 150.0,
			Weight: 1200.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "Boeing Defense",
			Country: "США",
			Website: "https://www.boeing.com/defense",
		},
		Tags:      []string{"топливный бак", "керосин", "алюминиевый сплав"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 10, 11, 20, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 10, 11, 20, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"объем_литры":              5000.0,
					"рабочее_давление_бар":     3.0,
					"температура_эксплуатации": -40.0,
				},
			},
		},
	},
	{
		Uuid:          "fuel-002",
		Name:          "Кислородный бак LOX",
		Description:   "Криогенный бак для жидкого кислорода. Оснащен многослойной теплоизоляцией.",
		Price:         950000.0,
		StockQuantity: 10,
		Category:      inventory_v1.Category_CATEGORY_FUEL,
		Dimensions: &inventory_v1.Dimensions{
			Length: 280.0,
			Width:  180.0,
			Height: 140.0,
			Weight: 900.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "Airbus Defence and Space",
			Country: "Германия",
			Website: "https://www.airbus.com/defence",
		},
		Tags:      []string{"кислородный бак", "криогенный", "теплоизоляция"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 25, 16, 30, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 25, 16, 30, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"объем_литры":          4000.0,
					"температура_хранения": -183.0,
					"толщина_изоляции_мм":  50.0,
				},
			},
		},
	},
	{
		Uuid:          "fuel-003",
		Name:          "Топливопровод высокого давления",
		Description:   "Топливопровод из титанового сплава для подачи топлива под высоким давлением.",
		Price:         150000.0,
		StockQuantity: 25,
		Category:      inventory_v1.Category_CATEGORY_FUEL,
		Dimensions: &inventory_v1.Dimensions{
			Length: 500.0,
			Width:  50.0,
			Height: 50.0,
			Weight: 80.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "Aerojet Rocketdyne",
			Country: "США",
			Website: "https://www.rocket.com",
		},
		Tags:      []string{"топливопровод", "титан", "высокое давление"},
		CreatedAt: timestamppb.New(time.Date(2024, 2, 1, 13, 45, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 2, 1, 13, 45, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"диаметр_мм":           50.0,
					"рабочее_давление_бар": 300.0,
					"толщина_стенки_мм":    3.0,
				},
			},
		},
	},

	// CATEGORY_PORT_HOLE - Детали для отверстий
	{
		Uuid:          "port-001",
		Name:          "Стыковочный узел APAS-95",
		Description:   "Андрогинный периферийный агрегат стыковки для космических кораблей. Совместим с МКС.",
		Price:         2200000.0,
		StockQuantity: 6,
		Category:      inventory_v1.Category_CATEGORY_PORT_HOLE,
		Dimensions: &inventory_v1.Dimensions{
			Length: 100.0,
			Width:  100.0,
			Height: 80.0,
			Weight: 250.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "РКК Энергия",
			Country: "Россия",
			Website: "https://www.energia.ru",
		},
		Tags:      []string{"стыковка", "МКС", "андрогинный"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 12, 8, 15, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 12, 8, 15, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"диаметр_люка_мм":       800.0,
					"максимальная_нагрузка": 10000.0,
					"герметичность_торр":    1e-6,
				},
			},
		},
	},
	{
		Uuid:          "port-002",
		Name:          "Люк выхода в открытый космос",
		Description:   "Герметичный люк для выхода космонавтов в открытый космос. Оснащен системой безопасности.",
		Price:         1800000.0,
		StockQuantity: 8,
		Category:      inventory_v1.Category_CATEGORY_PORT_HOLE,
		Dimensions: &inventory_v1.Dimensions{
			Length: 120.0,
			Width:  80.0,
			Height: 60.0,
			Weight: 180.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "Thales Alenia Space",
			Country: "Италия",
			Website: "https://www.thalesaleniaspace.com",
		},
		Tags:      []string{"люк", "выход в космос", "герметичность"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 18, 15, 20, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 18, 15, 20, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"диаметр_проема_мм":    650.0,
					"толщина_люка_мм":      25.0,
					"рабочее_давление_бар": 1.0,
				},
			},
		},
	},
	{
		Uuid:          "port-003",
		Name:          "Вентиляционный канал",
		Description:   "Вентиляционный канал для циркуляции воздуха в жилых отсеках космического корабля.",
		Price:         75000.0,
		StockQuantity: 30,
		Category:      inventory_v1.Category_CATEGORY_PORT_HOLE,
		Dimensions: &inventory_v1.Dimensions{
			Length: 200.0,
			Width:  30.0,
			Height: 30.0,
			Weight: 15.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "Lockheed Martin",
			Country: "США",
			Website: "https://www.lockheedmartin.com",
		},
		Tags:      []string{"вентиляция", "воздуховод", "жизнеобеспечение"},
		CreatedAt: timestamppb.New(time.Date(2024, 2, 8, 12, 10, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 2, 8, 12, 10, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"диаметр_мм":          300.0,
					"скорость_воздуха_мс": 2.0,
					"шум_дб":              45.0,
				},
			},
		},
	},

	// CATEGORY_WING - Детали для крыльев
	{
		Uuid:          "wing-001",
		Name:          "Солнечная панель ISS",
		Description:   "Солнечная панель для Международной космической станции. Высокоэффективные фотоэлементы.",
		Price:         3500000.0,
		StockQuantity: 4,
		Category:      inventory_v1.Category_CATEGORY_WING,
		Dimensions: &inventory_v1.Dimensions{
			Length: 800.0,
			Width:  300.0,
			Height: 50.0,
			Weight: 1200.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "Boeing Space Systems",
			Country: "США",
			Website: "https://www.boeing.com/space",
		},
		Tags:      []string{"солнечная панель", "энергетика", "фотоэлементы"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 5, 9, 30, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 5, 9, 30, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"мощность_квт":          32.0,
					"эффективность_процент": 28.5,
					"площадь_кв_м":          240.0,
				},
			},
		},
	},
	{
		Uuid:          "wing-002",
		Name:          "Аэродинамический стабилизатор",
		Description:   "Аэродинамический стабилизатор для спускаемого аппарата. Обеспечивает стабильность при спуске.",
		Price:         950000.0,
		StockQuantity: 12,
		Category:      inventory_v1.Category_CATEGORY_WING,
		Dimensions: &inventory_v1.Dimensions{
			Length: 150.0,
			Width:  80.0,
			Height: 20.0,
			Weight: 120.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "JAXA",
			Country: "Япония",
			Website: "https://www.jaxa.jp",
		},
		Tags:      []string{"стабилизатор", "аэродинамика", "спуск"},
		CreatedAt: timestamppb.New(time.Date(2024, 1, 22, 11, 45, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 1, 22, 11, 45, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"площадь_поверхности":   12.0,
					"угол_атаки_градусы":    15.0,
					"максимальная_нагрузка": 5000.0,
				},
			},
		},
	},
	{
		Uuid:          "wing-003",
		Name:          "Теплозащитный экран",
		Description:   "Теплозащитный экран из углерод-углеродного композита для защиты от высоких температур.",
		Price:         2800000.0,
		StockQuantity: 6,
		Category:      inventory_v1.Category_CATEGORY_WING,
		Dimensions: &inventory_v1.Dimensions{
			Length: 200.0,
			Width:  150.0,
			Height: 30.0,
			Weight: 400.0,
		},
		Manufacturer: &inventory_v1.Manufacturer{
			Name:    "NASA Ames Research Center",
			Country: "США",
			Website: "https://www.nasa.gov/ames",
		},
		Tags:      []string{"теплозащита", "углерод-углерод", "высокие температуры"},
		CreatedAt: timestamppb.New(time.Date(2024, 2, 12, 14, 20, 0, 0, time.UTC)),
		UpdatedAt: timestamppb.New(time.Date(2024, 2, 12, 14, 20, 0, 0, time.UTC)),
		Metadata: &inventory_v1.Part_MetadataDouble{
			MetadataDouble: &inventory_v1.DoubleMetadata{
				Data: map[string]float64{
					"максимальная_температура": 1650.0,
					"толщина_мм":               30.0,
					"плотность_кг_куб_м":       1600.0,
				},
			},
		},
	},
}
