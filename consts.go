package efrsb

const (
	prodURL = "https://bank-publications-prod.fedresurs.ru"
	devURL  = "https://bank-publications-demo.fedresurs.ru"
)

const (
	asRFC3339 = "2006-01-02T15:04:05"
)

// Типы отчетов
const (
	// ReportFinal Финальный отчет
	ReportFinal = "Final"
	// ReportAnnulment Аннулирование ранее опубликованного отчета
	ReportAnnulment = "Annulment"
	// ReportSignificantEvent Отчет по существенным фактам (устаревший)
	ReportSignificantEvent = "SignificantEvent"
	// ReportPeriodic Периодический отчет (устаревший)
	ReportPeriodic = "Periodic"
	// ReportFinal2 Финальный отчет
	ReportFinal2 = "Final2"
	// ReportAnnulment2 Аннулирование ранее опубликованного отчета
	ReportAnnulment2 = "Annulment2"
)

// Процедуры
const (
	// ProcedureFinancialRecovery Финансовое оздоровление
	ProcedureFinancialRecovery = "FinancialRecovery"
	// ProcedureExternalManagement Внешнее управление
	ProcedureExternalManagement = "ExternalManagement"
	// ProcedureTender Конкурсное производство
	ProcedureTender = "Tender"
	// ProcedureWatching Наблюдение
	ProcedureWatching = "Watching"
	// ProcedureCitizenAssetsDisposal Реализация имущества гражданина
	ProcedureCitizenAssetsDisposal = "CitizenAssetsDisposal"
	// ProcedureCitizenDebtRestructuring Реструктуризация долгов гражданина
	ProcedureCitizenDebtRestructuring = "CitizenDebtRestructuring"
)

// Типы сообщений
const (
	// MessageArbitralDecree Сообщение о судебном акте
	MessageArbitralDecree = "ArbitralDecree"
	// MessageAuction Объявление о проведении торгов
	MessageAuction = "Auction"
	// MessageMeeting Сообщение о собрании кредиторов
	MessageMeeting = "Meeting"
	// MessageMeetingResult Сообщение о результатах проведения собрания кредиторов
	MessageMeetingResult = "MeetingResult"
	// MessageTradeResult Сообщение о результатах торгов
	MessageTradeResult = "TradeResult"
	// MessageOther Иное сообщение
	MessageOther = "Other"
	// MessageAppointAdministration Решение о назначении временной администрации
	MessageAppointAdministration = "AppointAdministration"
	// MessageChangeAdministration Изменение состава временной администрации
	MessageChangeAdministration = "ChangeAdministration"
	// MessageTerminationAdministration Прекращение деятельности временной администрации
	MessageTerminationAdministration = "TerminationAdministration"
	//Message Annul Аннулирование ранее опубликованного сообщения
	MessageAnnul = "Annul"
	// MessageDemandAnnouncement Извещение о возможности предъявления требований
	MessageDemandAnnouncement = "DemandAnnouncement"
	// MessageCourtAssertAcceptance Объявление о принятии арбитражным судом заявления
	MessageCourtAssertAcceptance = "CourtAssertAcceptance"
	// MessageFinancialStateInformation Информация о финансовом состоянии
	MessageFinancialStateInformation = "FinancialStateInformation"
	// MessageBankPayment Объявление о выплатах Банка России
	MessageBankPayment = "BankPayment"
	// MessageAssetsReturning Объявление о возврате ценных бумаг и иного имущества
	MessageAssetsReturning = "AssetsReturning"
	// MessagePropertyInventoryResult Сведения о результатах инвентаризации имущества должника
	MessagePropertyInventoryResult = "PropertyInventoryResult"
	// MessagePropertyEvaluationReport Отчет оценщика об оценке имущества должника
	MessagePropertyEvaluationReport = "PropertyEvaluationReport"
	// MessageSaleContractResult Сведения о заключении договора купли-продажи
	MessageSaleContractResult = "SaleContractResult"
	// MessageCommittee Уведомление о проведении комитета кредиторов
	MessageCommittee = "Committee"
	// MessageCommitteeResult Сообщение о результатах проведения комитета кредиторов
	MessageCommitteeResult = "CommitteeResult"
	// MessageSaleOrderPledgedProperty	Об определении начальной продажной цены, утверждении порядка и условий проведения торгов по реализации предмета залога, порядка и условий обеспечения сохранности предмета залога
	MessageSaleOrderPledgedProperty = "SaleOrderPledgedProperty"
	// MessageReceivingCreditorDemand Уведомление о получении требований кредитора -
	MessageReceivingCreditorDemand = "ReceivingCreditorDemand"
	// MessageDeliberateBankruptcy Сообщение о наличии или об отсутствии признаков преднамеренного или фиктивного банкротства
	MessageDeliberateBankruptcy = "DeliberateBankruptcy"
	// MessageIntentionCreditOrg Сообщение о намерении исполнить обязательства кредитной организации
	MessageIntentionCreditOrg = "IntentionCreditOrg"
	// MessageLiabilitiesCreditOrg Сообщение о признании исполнения заявителем обязательств кредитной организации несостоявшимся
	MessageLiabilitiesCreditOrg = "LiabilitiesCreditOrg"
	// MessagePerformanceCreditOrg Сообщение об исполнении обязательств кредитной организации
	MessagePerformanceCreditOrg = "PerformanceCreditOrg"
	// MessageBuyingProperty Сообщение о преимущественном праве выкупа имущества
	MessageBuyingProperty = "BuyingProperty"
	// MessageDeclarationPersonDamages Заявление о привлечении контролирующих должника лиц, а также иных лиц, к ответственности в виде возмещения убытков
	MessageDeclarationPersonDamages = "DeclarationPersonDamages"
	// MessageActPersonDamages Судебный акт по результатам рассмотрения заявления о привлечении контролирующих должника лиц, а также иных лиц, к ответственности в виде возмещения убытков
	MessageActPersonDamages = "ActPersonDamages"
	// MessageActReviewPersonDamages Судебный акт по результатам пересмотра рассмотрения заявления о привлечении контролирующих должника лиц, а также иных лиц, к ответственности в виде возмещения убытков
	MessageActReviewPersonDamages = "ActReviewPersonDamages"
	// MessageDealInvalid Заявление о признании сделки должника недействительной
	MessageDealInvalid = "DealInvalid"
	// MessageActDealInvalid Судебный акт по результатам рассмотрения заявления об оспаривании сделки должника
	MessageActDealInvalid = "ActDealInvalid"
	// MessageActReviewDealInvalid Судебный акт по результатам пересмотра рассмотрения заявления об оспаривании сделки должника
	MessageActReviewDealInvalid = "ActReviewDealInvalid"
	// MessageDeclarationPersonSubsidiary Заявление о привлечении контролирующих должника лиц к субсидиарной ответственности
	MessageDeclarationPersonSubsidiary = "DeclarationPersonSubsidiary"
	// MessageActPersonSubsidiary Судебный акт по результатам рассмотрения заявления о привлечении контролирующих должника лиц к субсидиарной ответственности
	MessageActPersonSubsidiary = "ActPersonSubsidiary"
	// MessageActReviewPersonSubsidiary Судебный акт по результатам пересмотра рассмотрения заявления о привлечении контролирующих должника лиц к субсидиарной ответственности
	MessageActReviewPersonSubsidiary = "ActReviewPersonSubsidiary"
	// MessageMeetingWorker Уведомление о проведении собрания работников, бывших работников должника
	MessageMeetingWorker = "MeetingWorker"
	// MessageMeetingWorkerResult Сведения о решениях, принятых собранием работников, бывших работников должника
	MessageMeetingWorkerResult = "MeetingWorkerResult"
	// MessageViewDraftRestructuringPlan Сведения о порядке и месте ознакомления с проектом плана реструктуризации
	MessageViewDraftRestructuringPlan = "ViewDraftRestructuringPlan"
	// MessageViewExecRestructuringPlan Сведения о порядке и месте ознакомления с отчетом о результатах исполнения плана реструктуризации
	MessageViewExecRestructuringPlan = "ViewExecRestructuringPlan"
	// MessageChangeAuction Сообщение об изменении объявления о проведении торгов
	MessageChangeAuction = "ChangeAuction"
	// MessageCancelAuctionTradeResult Сообщение об отмене сообщения об объявлении торгов или сообщения о результатах торгов
	MessageCancelAuctionTradeResult = "CancelAuctionTradeResult"
	// MessageCancelDeliberateBankruptcy Сообщение об отмене сообщения о наличии или об отсутствии признаков преднамеренного или фиктивного банкротства
	MessageCancelDeliberateBankruptcy = "CancelDeliberateBankruptcy"
	// MessageChangeDeliberateBankruptcy Сообщение об изменении сообщения о наличии или об отсутствии признаков преднамеренного или фиктивного банкротства
	MessageChangeDeliberateBankruptcy = "ChangeDeliberateBankruptcy"
	// MessageTransferOwnershipRealEstate Сообщение о переходе права собственности на объект незавершенного строительства и прав на земельный участок
	MessageTransferOwnershipRealEstate = "TransferOwnershipRealEstate"
	// MessageSelectionPurchaserAssets Сведения о проведении отбора приобретателей имущества (активов) и обязательств кредитной организации
	MessageSelectionPurchaserAssets = "SelectionPurchaserAssets"
	// MessageEstimatesCurrentExpenses Сведения о смете текущих расходов кредитной организации или иной финансовой организации
	MessageEstimatesCurrentExpenses = "EstimatesCurrentExpenses"
	// MessageOrderAndTimingCalculations Сведения о порядке и сроках расчетов с кредиторами
	MessageOrderAndTimingCalculations = "OrderAndTimingCalculations"
	// MessageInformationAboutBankruptcy Информация о ходе конкурсного производства
	MessageInformationAboutBankruptcy = "InformationAboutBankruptcy"
	// MessageEstimatesAndUnsoldAssets Сведения об исполнении сметы текущих расходов и стоимости нереализованного имущества кредитной организации
	MessageEstimatesAndUnsoldAssets = "EstimatesAndUnsoldAssets"
	// MessageRemainingAssetsAndRight Объявление о наличии у кредитной организации оставшегося имущества и праве ее учредителей (участников) получить указанное имущество
	MessageRemainingAssetsAndRight = "RemainingAssetsAndRight"
	// MessageReducingSizeShareCapital Сообщение об уменьшении размера уставного капитала банка
	MessageReducingSizeShareCapital = "ReducingSizeShareCapital"
	// MessageImpendingTransferAssets Сообщение о предстоящей передаче приобретателю имущества (активов) и обязательств кредитной организации или их части
	MessageImpendingTransferAssets = "ImpendingTransferAssets"
	// MessageTransferAssets Сообщение о передаче приобретателю имущества и обязательств кредитной организации
	MessageTransferAssets = "TransferAssets"
	// MessageTransferInsurancePortfolio Уведомление о передаче страхового портфеля страховой организации
	MessageTransferInsurancePortfolio = "TransferInsurancePortfolio"
	// MessageSaleContractResult2 Сведения о заключении договора купли-продажи
	MessageSaleContractResult2 = "SaleContractResult2"
	// MessageBankOpenAccountDebtor Сведения о кредитной организации, в которой открыт специальный банковский счет должника
	MessageBankOpenAccountDebtor = "BankOpenAccountDebtor"
	// MessageRightUnsoldAsset Объявление о наличии непроданного имущества и праве собственника имущества должника – унитарного предприятия, учредителей (участников) должника получить такое имущество
	MessageRightUnsoldAsset = "RightUnsoldAsset"
	// MessageTransferResponsibilitiesFund Решение о передаче обязанности по выплате пожизненных негосударственных пенсий и средств пенсионных резервов другому негосударственному пенсионному фонду
	MessageTransferResponsibilitiesFund = "TransferResponsibilitiesFund"
	// MessageExtensionAdministration Продление срока деятельности временной администрации
	MessageExtensionAdministration = "ExtensionAdministration"
	// MessageProcedureGrantingIndemnity Предложение о погашении требований кредиторов путем предоставления отступного
	MessageProcedureGrantingIndemnity = "ProcedureGrantingIndemnity"
	// MessageStartSettlement Сообщение о начале расчетов
	MessageStartSettlement = "StartSettlement"
	// MessageProcessInventoryDebtor Сведения о ходе инвентаризации имущества должника
	MessageProcessInventoryDebtor = "ProcessInventoryDebtor"
	// MessageMeetingParticipantsBuilding Уведомление о проведении собрания участников строительства
	MessageMeetingParticipantsBuilding = "MeetingParticipantsBuilding"
	// MessageMeetingPartBuildResult Сообщение о результатах проведения собрания участников строительства
	MessageMeetingPartBuildResult = "MeetingPartBuildResult"
	// MessagePartBuildMonetaryClaim Извещение участникам строительства о возможности предъявления требований
	MessagePartBuildMonetaryClaim = "PartBuildMonetaryClaim"
	// MessageCourtAcceptanceStatement Сведения о принятии заявления о признании должника банкротом
	MessageCourtAcceptanceStatement = "CourtAcceptanceStatement"
	// MessageRebuttal Опровержение по решению суда опубликованных ранее сведений
	MessageRebuttal = "Rebuttal"
	// MessageActDealInvalid2 Судебный акт по результатам рассмотрения заявления об оспаривании сделки должника
	MessageActDealInvalid2 = "ActDealInvalid2"
	// MessageActReviewDealInvalid2 Судебный акт по результатам пересмотра рассмотрения заявления об оспаривании сделки должника
	MessageActReviewDealInvalid2 = "ActReviewDealInvalid2"
	// MessageAccessionDeclarationSubsidiary Предложение о присоединении к заявлению о привлечении контролирующих лиц должника к субсидиарной ответственности
	MessageAccessionDeclarationSubsidiary = "AccessionDeclarationSubsidiary"
	// MessageCreditorChoiceRightSubsidiary Сообщение о праве кредитора выбрать способ распоряжения правом требования о привлечении к субсидиарной ответственности
	MessageCreditorChoiceRightSubsidiary = "CreditorChoiceRightSubsidiary"
	// MessageDisqualificationArbitrationManager Сообщение о дисквалификации арбитражного управляющего
	MessageDisqualificationArbitrationManager = "DisqualificationArbitrationManager"
	// MessageChangeEstimatesCurrentExpenses Сведения о скорректированной смете текущих расходов кредитной организации или иной
	MessageChangeEstimatesCurrentExpenses = "ChangeEstimatesCurrentExpenses"
	// MessageDisqualificationArbitrationManager2 Сообщение о дисквалификации арбитражного управляющего
	MessageDisqualificationArbitrationManager2 = "DisqualificationArbitrationManager2"
	// MessageActPersonSubsidiary2 Судебный акт по результатам рассмотрения заявления о привлечении контролирующих лиц к субсидиарной ответственности
	MessageActPersonSubsidiary2 = "ActPersonSubsidiary2"
	// MessageActReviewPersonSubsidiary2 Судебный акт по результатам пересмотра рассмотрения заявления о привлечении контролирующих лиц к субсидиарной ответственности
	MessageActReviewPersonSubsidiary2 = "ActReviewPersonSubsidiary2"
	// MessageAssessmentReport Отчет оценщика об оценке имущества должника
	MessageAssessmentReport = "AssessmentReport"
	// MessageReturnOfApplicationOnExtrajudicialBankruptcy Сообщение о возврате гражданину поданного им заявления о признании гражданина банкротом во внесудебном порядке
	MessageReturnOfApplicationOnExtrajudicialBankruptcy = "ReturnOfApplicationOnExtrajudicialBankruptcy"
	// MessageStartOfExtrajudicialBankruptcy Сообщение о возбуждении процедуры внесудебного банкротства гражданина
	MessageStartOfExtrajudicialBankruptcy = "StartOfExtrajudicialBankruptcy"
	// MessageTerminationOfExtrajudicialBankruptcy Сообщение о прекращении процедуры внесудебного банкротства гражданина
	MessageTerminationOfExtrajudicialBankruptcy = "TerminationOfExtrajudicialBankruptcy"
	// MessageCompletionOfExtrajudicialBankruptcy Сообщение о завершении процедуры внесудебного банкротства гражданина
	MessageCompletionOfExtrajudicialBankruptcy = "CompletionOfExtrajudicialBankruptcy"
	// MessageCommittee2 Уведомление о проведении комитета кредиторов
	MessageCommittee2 = "Committee2"
	// MessageMeeting2 Сообщение о собрании кредиторов
	MessageMeeting2 = "Meeting2"
	// MessageMeetingWorker2 Уведомление о проведении собрания работников, бывших работников должника
	MessageMeetingWorker2 = "MeetingWorker2"
	// MessageMeetingParticipantsBuilding2 Уведомление о проведении собрания участников строительства
	MessageMeetingParticipantsBuilding2 = "MeetingParticipantsBuilding2"
	// MessageAuction2 Объявление о проведении торгов
	MessageAuction2 = "Auction2"
	// MessageChangeAuction2 Сообщение об изменении объявления о проведении торгов
	MessageChangeAuction2 = "ChangeAuction2"
	// MessageSaleOrderPledgedProperty2 Об определении начальной продажной цены, утверждении порядка и условий проведения торгов по реализации предмета залога, порядка и условий обеспечения сохранности предмета залога
	MessageSaleOrderPledgedProperty2 = "SaleOrderPledgedProperty2"
	// MessageExtraordinaryExpenses Сведения об осуществлении внеочередных расходов
	MessageExtraordinaryExpenses = "ExtraordinaryExpenses"
	// MessageReturnOfApplicationOnExtrajudicialBankruptcy2 Сообщение о возврате гражданину поданного им заявления о признании гражданина банкротом во внесудебном порядке
	MessageReturnOfApplicationOnExtrajudicialBankruptcy2 = "ReturnOfApplicationOnExtrajudicialBankruptcy2"
	// MessageCourseOfSalePersonProperty Сведения об утверждении положения о порядке, об условиях и о сроках реализации имущества гражданина
	MessageCourseOfSalePersonProperty = "CourseOfSalePersonProperty"
	// MessageApplicationReviewCourtDecision Сведения о подаче заявления о пересмотре судебного акта по вновь открывшимся обстоятельствам
	MessageApplicationReviewCourtDecision = "ApplicationReviewCourtDecision"
	// MessageChangeCreditorChoiceRightSubsidiary Сообщение о праве изменить способ распоряжения правом требования о привлечении к субсидиарной ответственности
	MessageChangeCreditorChoiceRightSubsidiary = "ChangeCreditorChoiceRightSubsidiary"
	// MessageRightOfAcquisitionExecution Сообщение о реализации права приобрести право требования о привлечении к субсидиарной ответственности
	MessageRightOfAcquisitionExecution = "RightOfAcquisitionExecution"
	// MessageRefusalOfIntentionCreditOrg Сообщение о реализации права приобрести право требования о привлечении к субсидиарной ответственности
	MessageRefusalOfIntentionCreditOrg = "RefusalOfIntentionCreditOrg"
	// MessageLiabilitiesCreditOrgAct Определение о признании исполнения заявителем обязательств несостоявшимся
	MessageLiabilitiesCreditOrgAct = "LiabilitiesCreditOrgAct"
	// MessageCreditorsDemandRegistered Сообщение о включении заявленных требований в реестр требований кредиторов
	MessageCreditorsDemandRegistered = "CreditorsDemandRegistered"
	// MessageIntentionOfDemandsFulfilment Сведения о вынесении судебного акта об удовлетворении заявления о намерении удовлетворить в полном объеме требования кредиторов к должнику
	MessageIntentionOfDemandsFulfilment = "IntentionOfDemandsFulfilment"
	// MessageIntentionOfDemandsFulfilmentReview Сведения о пересмотре судебного акта об удовлетворении заявления о намерении удовлетворить в полном объеме требования кредиторов к должнику
	MessageIntentionOfDemandsFulfilmentReview = "IntentionOfDemandsFulfilmentReview"
	// MessageFulfilledDemandsRecognition Сведения о вынесении судебного акта о признании требований кредиторов удовлетворенными
	MessageFulfilledDemandsRecognition = "FulfilledDemandsRecognition"
	// MessageFulfilledDemandsRecognitionReview Сведения о пересмотре судебного акта о признании требований кредиторов удовлетворенными
	MessageFulfilledDemandsRecognitionReview = "FulfilledDemandsRecognitionReview"
	// MessageDealInvalidResult Результаты рассмотрения/пересмотра заявлений об оспаривании сделок
	MessageDealInvalidResult = "DealInvalidResult"
	// MessageReceivingCreditorDemand2 Уведомление о получении требований кредитора
	MessageReceivingCreditorDemand2 = "ReceivingCreditorDemand2"
	// MessageDealInvalid2 Заявление о признании сделки должника недействительной
	MessageDealInvalid2 = "DealInvalid2"
	// MessageDealInvalidResult2 Результаты рассмотрения/пересмотра заявлений об оспаривании сделок
	MessageDealInvalidResult2 = "DealInvalidResult2"
	// MessagePersonResponsibilityDeclaration Заявление о привлечении контролирующих должника и иных лиц к ответственности
	MessagePersonResponsibilityDeclaration = "PersonResponsibilityDeclaration"
	// MessagePersonResponsibilityResult Результаты рассмотрения/пересмотра заявлений о привлечении контролирующих должника и иных лиц к ответственности
	MessagePersonResponsibilityResult = "PersonResponsibilityResult"
)

// Типы актов
const (
	// ActExternalManagement о введении внешнего управления
	ActExternalManagement = "ExternalManagement"
	// ActChangeCourt об отмене или изменении судебных актов
	ActChangeCourt = "ChangeCourtAct"
	// ActLegalCaseResume о возобновлении производства по делу о несостоятельности (банкротстве)
	ActLegalCaseResume = "LegalCaseResume"
	// ActArbitrManagerApproval об утверждении арбитражного управляющего
	ActArbitrManagerApproval = "ArbitrManagerApproval"
	// ActArbitrManagerRelease об освобождении или отстранении арбитражного управляющего
	ActArbitrManagerRelease = "ArbitrManagerRelease"
	// ActReceivership о признании должника банкротом и открытии конкурсного производства
	ActReceivership = "Receivership"
	// ActLegalCaseTermination о прекращении производства по делу
	ActLegalCaseTermination = "LegalCaseTermination"
	// ActFinancialRecovery о введении финансового оздоровления
	ActFinancialRecovery = "FinancialRecovery"
	// ActBankruptcyRefusal об отказе в признании должника банкротом
	ActBankruptcyRefusal = "BankruptcyRefusal"
	// ActObservation о введении наблюдения
	ActObservation = "Observation"
	// ActOther Другие судебные акты
	ActOther = "OtherAct"
	// ActOtherDefinition Другие определения
	ActOtherDefinition = "OtherDefinition"
	// ActThirdPartyPayoffRequestApproval об удовлетворении заявлений третьих лиц о намерении погасить обязательства должника
	ActThirdPartyPayoffRequestApproval = "ThirdPartyPayoffRequestApproval"
	// ActDebtRestructuring о признании обоснованным заявления о признании гражданина банкротом и введении реструктуризации его долгов
	ActDebtRestructuring = "DebtRestructuring"
	// ActPropertySale о признании гражданина банкротом и введении реализации имущества гражданина
	ActPropertySale = "PropertySale"
	// ActDebtRestructuringPlan об утверждении плана реструктуризации долгов гражданина
	ActDebtRestructuringPlan = "DebtRestructuringPlan"
	// ActDebtRestructuringComplete о завершении реструктуризации долгов гражданина
	ActDebtRestructuringComplete = "DebtRestructuringComplete"
	// ActArbitrManagerActionsIllegal о признании действий (бездействий) арбитражного управляющего незаконными
	ActArbitrManagerActionsIllegal = "ArbitrManagerActionsIllegal"
	// ActArbitrManagerLossesRecovery о взыскании с арбитражного управляющего убытков в связи с неисполнением или ненадлежащим исполнением обязанностей
	ActArbitrManagerLossesRecovery = "ArbitrManagerLossesRecovery"
	// ActObligationsDischargeRefusal о неприменении в отношении гражданина правила об освобождении от исполнения обязательств
	ActObligationsDischargeRefusal = "ObligationsDischargeRefusal"
	// ActPropertySaleComplete о завершении реализации имущества гражданина
	ActPropertySaleComplete = "PropertySaleComplete"
	// ActDeveloperBankruptcy о применении при банкротстве должника правил параграфа «Банкротство застройщиков»
	ActDeveloperBankruptcy = "DeveloperBankruptcy"
	// ActCaseTransfer о передаче дела на рассмотрение другого арбитражного суда
	ActCaseTransfer = "CaseTransfer"
	// ActLegalCaseEnd о завершении конкурсного производства
	ActLegalCaseEnd = "LegalCaseEnd"
	// ActExtensionProcedure о продлении срока процедуры
	ActExtensionProcedure = "ExtensionProcedure"
	// ActChangeArbitralDecree об изменении судебного акта
	ActChangeArbitralDecree = "ChangeArbitralDecree"
	// ActCancelArbitralDecree об отмене судебного акта
	ActCancelArbitralDecree = "CancelArbitralDecree"
)

// Типы сообщений ЭТП
const (
	// EtpBiddingInvitation Сообщение о продаже
	EtpBiddingInvitation = "BiddingInvitation"
	// EtpBiddingDeclaration Объявлены торги
	EtpBiddingDeclaration = "BiddingDeclaration"
	// EtpApplicationSessionStart Начат прием заявок
	EtpApplicationSessionStart = "ApplicationSessionStart"
	// EtpApplicationSessionEnd Приём заявок закончен
	EtpApplicationSessionEnd = "ApplicationSessionEnd"
	// EtpApplicationSessionStatistic Сведения о ходе проведения торгов
	EtpApplicationSessionStatistic = "ApplicationSessionStatistic"
	// EtpBiddingStart Начаты торги
	EtpBiddingStart = "BiddingStart"
	// EtpBiddingProcess Предложение о цене
	EtpBiddingProcess = "BiddingProcess"
	// EtpBiddingEnd Торги завершены
	EtpBiddingEnd = "BiddingEnd"
	// EtpBiddingResult Результаты торгов
	EtpBiddingResult = "BiddingResult"
	// EtpSaleAgreement Сведения о заключении договора купли-продажи
	EtpSaleAgreement = "SaleAgreement"
	//EtpAnnulment  Аннулирование
	EtpAnnulment = "Annulment"
	// EtpBiddingCancel Отмена торгов
	EtpBiddingCancel = "BiddingCancel"
	// EtpBiddingFail Торги не состоялись
	EtpBiddingFail = "BiddingFail"
	// EtpBiddingPause Торги приостановлены
	EtpBiddingPause = "BiddingPause"
	// EtpBiddingResume Торги возобновлены
	EtpBiddingResume = "BiddingResume"
	// EtpContractSale Сообщение о заключении договора купли-продажи
	EtpContractSale = "ContractSale"
	// EtpBiddingEndBankruptcyCreditor О завершении торгов вследствие оставления конкурсным кредитором предмета залога за собой
	EtpBiddingEndBankruptcyCreditor = "BiddingEndBankruptcyCreditor"
	// EtpErrorMessage Сообщение о возникновении технического сбоя
	EtpErrorMessage = "ErrorMessage"
	// EtpBiddingNewTerm Сообщение об установлении новых сроков
	EtpBiddingNewTerm = "BiddingNewTerm"
)
