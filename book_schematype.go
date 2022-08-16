package openbd

// 扱い社コード
type agentIdentifierItem struct {
	AgentIdtype string `json:"AgentIDType,omitempty"`
	IdtypeName  string `json:"IDTypeName,omitempty"`
	Idvalue     string `json:"IDValue,omitempty"`
}

// 商品の対象読者を示す方法
type audienceItem struct {
	AudienceCodeType  string `json:"AudienceCodeType,omitempty"`
	AudienceCodeValue string `json:"AudienceCodeValue,omitempty"`
}

// 著者補足情報
type authorItem struct {
	Dokujikubun string  `json:"dokujikubun"`
	Listseq     float64 `json:"listseq"`
}

// 商品の付帯項目
type collateralDetail struct {
	SupportingResource []supportingResourceItem `json:"SupportingResource,omitempty"`
	TextContent        []textContentItem        `json:"TextContent,omitempty"`
}

// シリーズに関する情報
type collection struct {
	CollectionSequence      collectionSequence            `json:"CollectionSequence,omitempty"`
	CollectionSequenceArray []collectionSequenceArrayItem `json:"CollectionSequenceArray,omitempty"`
	CollectionType          string                        `json:"CollectionType,omitempty"`
	TitleDetail             collectionTitleDetail         `json:"TitleDetail,omitempty"`
}

// 「配本回数」について（複数に変更されたため互換性のために維持→CollectionSequenceArrayを参照してください）
type collectionSequence struct {
	CollectionSequenceNumber   string `json:"CollectionSequenceNumber,omitempty"`
	CollectionSequenceType     string `json:"CollectionSequenceType,omitempty"`
	CollectionSequenceTypeName string `json:"CollectionSequenceTypeName,omitempty"`
}

// 「配本回数」について
type collectionSequenceArrayItem struct {
	CollectionSequenceNumber   string `json:"CollectionSequenceNumber,omitempty"`
	CollectionSequenceType     string `json:"CollectionSequenceType,omitempty"`
	CollectionSequenceTypeName string `json:"CollectionSequenceTypeName,omitempty"`
}

// シリーズタイトル
type collectionTitleDetail struct {
	TitleElement []titleElementItem `json:"TitleElement,omitempty"`
	TitleType    string             `json:"TitleType,omitempty"`
}

// 商品の著者情報
type contributorItem struct {
	BiographicalNote string                `json:"BiographicalNote,omitempty"`
	ContributorRole  []contributorRoleItem `json:"ContributorRole,omitempty"`
	PersonName       personName            `json:"PersonName,omitempty"`
	SequenceNumber   string                `json:"SequenceNumber,omitempty"`
}

// 著者区分
type contributorRoleItem string

// 商品に関する情報
type descriptiveDetail struct {
	Audience           []audienceItem               `json:"Audience,omitempty"`
	Collection         collection                   `json:"Collection,omitempty"`
	Contributor        []contributorItem            `json:"Contributor,omitempty"`
	EditionStatement   string                       `json:"EditionStatement,omitempty"`
	EditionType        string                       `json:"EditionType,omitempty"`
	Extent             []extentItem                 `json:"Extent,omitempty"`
	Language           []languageItem               `json:"Language,omitempty"`
	Measure            []measureItem                `json:"Measure,omitempty"`
	ProductComposition string                       `json:"ProductComposition,omitempty"`
	ProductForm        string                       `json:"ProductForm,omitempty"`
	ProductFormDetail  string                       `json:"ProductFormDetail,omitempty"`
	ProductPart        []productPartItem            `json:"ProductPart,omitempty"`
	Subject            []subjectItem                `json:"Subject,omitempty"`
	TitleDetail        descriptiveDetailTitleDetail `json:"TitleDetail,omitempty"`
}

// 商品の「書名」を設定
type descriptiveDetailTitleDetail struct {
	TitleElement titleElement `json:"TitleElement,omitempty"`
	TitleType    string       `json:"TitleType,omitempty"`
}

// 商品に関連する（数値的）範囲、程度など
type extentItem struct {
	ExtentType  string `json:"ExtentType,omitempty"`
	ExtentUnit  string `json:"ExtentUnit,omitempty"`
	ExtentValue string `json:"ExtentValue,omitempty"`
}

// 版元独自書誌
type hanmoto struct {
	Author                 []authorItem `json:"author,omitempty"`
	Bessoushiryou          string       `json:"bessoushiryou,omitempty"`
	Bikoujpo               string       `json:"bikoujpo,omitempty"`
	Bikoutrc               string       `json:"bikoutrc,omitempty"`
	Datecreated            Date         `json:"datecreated,omitempty"`
	Datejuuhanyotei        string       `json:"datejuuhanyotei,omitempty"`
	Datekoukai             Date         `json:"datekoukai,omitempty"`
	Datemodified           Date         `json:"datemodified"`
	Dateshuppan            string       `json:"dateshuppan,omitempty"`
	Datezeppan             string       `json:"datezeppan,omitempty"`
	Dokushakakikomi        string       `json:"dokushakakikomi,omitempty"`
	Dokushakakikomipagesuu float64      `json:"dokushakakikomipagesuu,omitempty"`
	Furoku                 float64      `json:"furoku,omitempty"`
	Furokusonota           string       `json:"furokusonota,omitempty"`
	Gatsugougousuu         string       `json:"gatsugougousuu,omitempty"`
	Genrecodetrc           float64      `json:"genrecodetrc,omitempty"`
	Genrecodetrcjidou      float64      `json:"genrecodetrcjidou,omitempty"`
	Genshomei              string       `json:"genshomei,omitempty"`
	Han                    string       `json:"han,omitempty"`
	Hanmotoinfo            hanmotoinfo  `json:"hanmotoinfo,omitempty"`
	Hanmotokarahitokoto    string       `json:"hanmotokarahitokoto,omitempty"`
	Hastameshiyomi         bool         `json:"hastameshiyomi,omitempty"`
	Hatsubai               string       `json:"hatsubai,omitempty"`
	Hatsubaiyomi           string       `json:"hatsubaiyomi,omitempty"`
	Jushoujouhou           string       `json:"jushoujouhou,omitempty"`
	Jyuhan                 []jyuhanItem `json:"jyuhan,omitempty"`
	Kaisetsu105w           string       `json:"kaisetsu105w,omitempty"`
	Kankoukeitai           string       `json:"kankoukeitai,omitempty"`
	Kanrensho              string       `json:"kanrensho,omitempty"`
	Kanrenshoisbn          string       `json:"kanrenshoisbn,omitempty"`
	Kubunhanbai            bool         `json:"kubunhanbai,omitempty"`
	Lanove                 bool         `json:"lanove,omitempty"`
	Maegakinado            string       `json:"maegakinado,omitempty"`
	Ndccode                string       `json:"ndccode,omitempty"`
	Obinaiyou              string       `json:"obinaiyou,omitempty"`
	Reviews                []review     `json:"reviews,omitempty"`
	Rubynoumu              bool         `json:"rubynoumu,omitempty"`
	Ruishokyougousho       string       `json:"ruishokyougousho,omitempty"`
	Sonotatokkijikou       string       `json:"sonotatokkijikou,omitempty"`
	Storelink              string       `json:"storelink,omitempty"`
	Toji                   string       `json:"toji,omitempty"`
	Tsuiki                 string       `json:"tsuiki,omitempty"`
	Zaiko                  float64      `json:"zaiko,omitempty"`
	Zasshicode             string       `json:"zasshicode,omitempty"`
}

// 版元情報
type hanmotoinfo struct {
	Chokutori        string `json:"chokutori,omitempty"`
	Eigyoudaihyousha string `json:"eigyoudaihyousha,omitempty"`
	Emailshoten      string `json:"emailshoten,omitempty"`
	Facebook         string `json:"facebook,omitempty"`
	Facsimileshoten  string `json:"facsimileshoten,omitempty"`
	Henpin           string `json:"henpin,omitempty"`
	Jiyuukinyuu      string `json:"jiyuukinyuu,omitempty"`
	Name             string `json:"name,omitempty"`
	Ordersite        string `json:"ordersite,omitempty"`
	Ordersitejisha   string `json:"ordersitejisha,omitempty"`
	Ordersitesonota  string `json:"ordersitesonota,omitempty"`
	Phoneshoten      string `json:"phoneshoten,omitempty"`
	Toritsugisonota  string `json:"toritsugisonota,omitempty"`
	Toritsugitorikyo string `json:"toritsugitorikyo,omitempty"`
	Twitter          string `json:"twitter,omitempty"`
	URL              string `json:"url,omitempty"`
	Yomi             string `json:"yomi,omitempty"`
}

// 「発行元出版社」に関する情報
type imprint struct {
	ImprintIdentifier []imprintIdentifierItem `json:"ImprintIdentifier,omitempty"`
	ImprintName       string                  `json:"ImprintName,omitempty"`
}

// 「発行元出版社」コード
type imprintIdentifierItem struct {
	Idvalue       string `json:"IDValue,omitempty"`
	ImprintIdtype string `json:"ImprintIDType,omitempty"`
}

// 重版情報
type jyuhanItem struct {
	Comment string  `json:"comment,omitempty"`
	Ctime   Date    `json:"ctime,omitempty"`
	Date    Date    `json:"date"`
	Suri    float64 `json:"suri,omitempty"`
}

// 表記に利用する言語
type languageItem struct {
	CountryCode  string `json:"CountryCode,omitempty"`
	LanguageCode string `json:"LanguageCode,omitempty"`
	LanguageRole string `json:"LanguageRole,omitempty"`
}

// 扱い社に関する情報
type marketPublishingDetail struct {
	MarketPublishingStatus     string                        `json:"MarketPublishingStatus,omitempty"`
	MarketPublishingStatusNote string                        `json:"MarketPublishingStatusNote,omitempty"`
	PublisherRepresentative    []publisherRepresentativeItem `json:"PublisherRepresentative,omitempty"`
}

// タテ・ヨコの実寸
type measureItem struct {
	MeasureType     string `json:"MeasureType,omitempty"`
	MeasureUnitCode string `json:"MeasureUnitCode,omitempty"`
	Measurement     string `json:"Measurement,omitempty"`
}

type onix struct {
	CollateralDetail  collateralDetail      `json:"CollateralDetail,omitempty"`
	DescriptiveDetail descriptiveDetail     `json:"DescriptiveDetail,omitempty"`
	NotificationType  string                `json:"NotificationType,omitempty"`
	ProductIdentifier onixProductIdentifier `json:"ProductIdentifier,omitempty"`
	ProductSupply     productSupply         `json:"ProductSupply,omitempty"`
	PublishingDetail  publishingDetail      `json:"PublishingDetail,omitempty"`
	RecordReference   string                `json:"RecordReference,omitempty"`
	RelatedMaterial   []relatedMaterialItem `json:"RelatedMaterial,omitempty"`
}

// 識別情報(ISBN)
type onixProductIdentifier struct {
	Idvalue       string `json:"IDValue,omitempty"`
	ProductIdtype string `json:"ProductIDType,omitempty"`
}

// 著者名
type personName struct {
	Collationkey string `json:"collationkey,omitempty"`
	Content      string `json:"content,omitempty"`
}

// 特価期限
type priceDateItem struct {
	Date          string `json:"Date,omitempty"`
	Price         string `json:"Price,omitempty"`
	PriceDateRole string `json:"PriceDateRole,omitempty"`
}

// 価格（単価）について
type priceItem struct {
	CurrencyCode string          `json:"CurrencyCode,omitempty"`
	PriceAmount  string          `json:"PriceAmount,omitempty"`
	PriceDate    []priceDateItem `json:"PriceDate,omitempty"`
	PriceType    string          `json:"PriceType,omitempty"`
}

// 付録の有無
type productPartItem struct {
	NumberOfItemsOfThisForm string `json:"NumberOfItemsOfThisForm,omitempty"`
	ProductForm             string `json:"ProductForm,omitempty"`
	ProductFormDescription  string `json:"ProductFormDescription,omitempty"`
}

// 市場における商品の出版状況、供給等
type productSupply struct {
	MarketPublishingDetail marketPublishingDetail `json:"MarketPublishingDetail,omitempty"`
	SupplyDetail           supplyDetail           `json:"SupplyDetail,omitempty"`
}

// 「発行元出版社」と異なる場合、「発売元出版社」についての情報
type publisher struct {
	PublisherIdentifier []publisherIdentifierItem `json:"PublisherIdentifier,omitempty"`
	PublisherName       string                    `json:"PublisherName,omitempty"`
	PublishingRole      string                    `json:"PublishingRole,omitempty"`
}

// 「発売元出版社」コード
type publisherIdentifierItem struct {
	Idvalue         string `json:"IDValue,omitempty"`
	PublisherIdtype string `json:"PublisherIDType,omitempty"`
}

// 扱い社
type publisherRepresentativeItem struct {
	AgentIdentifier []agentIdentifierItem `json:"AgentIdentifier,omitempty"`
	AgentName       string                `json:"AgentName,omitempty"`
	AgentRole       string                `json:"AgentRole,omitempty"`
}

// 商品の出版に関する日付情報
type publishingDateItem struct {
	Date               string `json:"Date,omitempty"`
	PublishingDateRole string `json:"PublishingDateRole,omitempty"`
}

// 出版社に関する情報
type publishingDetail struct {
	Imprint          imprint              `json:"Imprint,omitempty"`
	Publisher        publisher            `json:"Publisher,omitempty"`
	PublishingDate   []publishingDateItem `json:"PublishingDate,omitempty"`
	PublishingStatus string               `json:"PublishingStatus,omitempty"`
}

// 旧版商品ISBN
type relatedMaterialItem struct {
	RelatedProduct relatedProduct `json:"RelatedProduct,omitempty"`
}

// 旧版の商品が存在する場合にその商品のISBNを設定する。
type relatedProduct struct {
	ProductIdentifier   relatedProductProductIdentifier `json:"ProductIdentifier,omitempty"`
	ProductRelationCode string                          `json:"ProductRelationCode,omitempty"`
}

type relatedProductProductIdentifier struct {
	Idvalue       string `json:"IDValue,omitempty"`
	ProductIdtype string `json:"ProductIDType,omitempty"`
}

// 補助リソースのバージョン
type resourceVersionItem struct {
	ResourceForm           string        `json:"ResourceForm,omitempty"`
	ResourceLink           string        `json:"ResourceLink,omitempty"`
	ResourceVersionFeature []interface{} `json:"ResourceVersionFeature,omitempty"`
}

// 返品条件について
type returnsConditions struct {
	ReturnsCode     string `json:"ReturnsCode,omitempty"`
	ReturnsCodeType string `json:"ReturnsCodeType,omitempty"`
}

// 書評情報
type review struct {
	Appearance string  `json:"appearance,omitempty"`
	Choyukan   string  `json:"choyukan,omitempty"`
	Gou        string  `json:"gou,omitempty"`
	Han        string  `json:"han,omitempty"`
	KubunID    float64 `json:"kubun_id,omitempty"`
	Link       string  `json:"link,omitempty"`
	PostUser   string  `json:"post_user,omitempty"`
	Reviewer   string  `json:"reviewer,omitempty"`
	Source     string  `json:"source,omitempty"`
	SourceID   float64 `json:"source_id,omitempty"`
}

type Book struct {
	Hanmoto hanmoto `json:"hanmoto,omitempty"`
	Onix    onix    `json:"onix,omitempty"`
	Summary summary `json:"summary,omitempty"`
}

// 商品の主題、テーマ、カテゴリなど
type subjectItem struct {
	MainSubject             string `json:"MainSubject,omitempty"`
	SubjectCode             string `json:"SubjectCode,omitempty"`
	SubjectHeadingText      string `json:"SubjectHeadingText,omitempty"`
	SubjectSchemeIdentifier string `json:"SubjectSchemeIdentifier,omitempty"`
}

// サブタイトル
type subtitle struct {
	Collationkey string `json:"collationkey,omitempty"`
	Content      string `json:"content,omitempty"`
}

// 書誌の概要
type summary struct {
	Author    string `json:"author"`
	Cover     string `json:"cover"`
	Isbn      string `json:"isbn"`
	Pubdate   string `json:"pubdate"`
	Publisher string `json:"publisher"`
	Series    string `json:"series"`
	Title     string `json:"title"`
	Volume    string `json:"volume"`
}

// 商品のアベイラビリテイ（入手可能性）や価格、その他の供給状況
type supplyDetail struct {
	Price               []priceItem       `json:"Price,omitempty"`
	ProductAvailability string            `json:"ProductAvailability,omitempty"`
	ReturnsConditions   returnsConditions `json:"ReturnsConditions,omitempty"`
}

// ONIXデータを補助する販促情報や追加情報などのデジタルデータ
type supportingResourceItem struct {
	ContentAudience     string                `json:"ContentAudience,omitempty"`
	ResourceContentType string                `json:"ResourceContentType,omitempty"`
	ResourceMode        string                `json:"ResourceMode,omitempty"`
	ResourceVersion     []resourceVersionItem `json:"ResourceVersion,omitempty"`
}

// 内容紹介
type textContentItem struct {
	ContentAudience string `json:"ContentAudience,omitempty"`
	Text            string `json:"Text,omitempty"`
	TextType        string `json:"TextType,omitempty"`
}

// タイトルの要素
type titleElement struct {
	PartNumber        string                `json:"PartNumber,omitempty"`
	Subtitle          subtitle              `json:"Subtitle,omitempty"`
	TitleElementLevel string                `json:"TitleElementLevel,omitempty"`
	TitleText         titleElementTitleText `json:"TitleText,omitempty"`
}

// シリーズ・レーベルの名前
type titleElementItem struct {
	PartNumber        string                    `json:"PartNumber,omitempty"`
	TitleElementLevel string                    `json:"TitleElementLevel,omitempty"`
	TitleText         titleElementItemTitleText `json:"TitleText,omitempty"`
}

// シリーズ名
type titleElementItemTitleText struct {
	Collationkey string `json:"collationkey,omitempty"`
	Content      string `json:"content,omitempty"`
}

// 書名
type titleElementTitleText struct {
	Collationkey string `json:"collationkey,omitempty"`
	Content      string `json:"content,omitempty"`
}
