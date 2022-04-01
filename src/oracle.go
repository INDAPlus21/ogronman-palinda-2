// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
	nouns  = "way year work government day man men world life lives part house houses course case system place end group company companies partyparties information informations school fact moneymoneys point example state business businesses night areaareas water thing family families head hand order john side home development week power country countries council use service room market problem problems court lot aaes war police interestinterests car lawlaws road form face education policypolicies researchresearches sort office bodybodies personpersons healthhealths mother question period name book level childchildren control societysocieties minister view door line communitycommunities south citycities god father centrecentres effect staffstaffs position kind job womanwomen action managementmanagements act processprocesses north age evidenceevidences ideaideas west support moment sense report mind churchchurches morning death change industryindustries land care centurycenturies range table back trade historyhistories studystudies street committee rate word foodfoods language experience result team other sirsirs section programme air authorityauthorities role reason price town classclasses nature subject department union bank member value need east practicepractices type paper date decision figure right wifewives president universityuniversities friend club qualityqualities voice lord stage king us situation light taxtaxes production marchmarches secretarysecretaries art board may hospital month musicmusics cost field award issue bed project chapter girl game amount basisbases knowledgeknowledges approachapproaches seriesseries love top news front future manager account computer securitysecurities rest labour structure hair bill heart force attention movement successsuccesses letter agreement capital analysisanalyses population environment performance model material theorytheories growth fire chance boy relationship son sea record size propertyproperties space term director plan behaviourbehaviours treatment energyenergies st peter income cup scheme design response association choice pressure hall couple technologytechnologies defence list chairmanchairmen losslosses activityactivities contract countycounties wall paul difference armyarmies hotel sun product summer set village colourcolours floor season unit park hour investment test garden husband employment style science look deal charge helphelps economyeconomies new page risk advice event picture commission fishfish college oil doctor opportunityopportunities film conference operation application presspresses extent addition station window shop accessaccesses region doubt majoritymajorities degree television blood statement sound election parliament site mark importance title speciesspecies increase return concern publicpublics competition software glassglasses ladyladies answer earth daughterdaughters purpose responsibilityresponsibilities leader river eyeeyes abilityabilities appeal opposition campaign respect task instance sale whole officer method division source piece pattern lack disease equipmentequipments surface oxford demand post mouth radio provision attempt sector firm statusstatuses peacepeaces varietyvarieties teacher showshows speaker babybabies arm base missmisses safetysafeties trouble culture direction context charactercharacters boxboxes discussion past weight organisationorganisations start brotherbrothers league condition machine argument sexsexes budget english transport share mum cashcashes principle exchange aid librarylibraries version rule tea balance afternoon reference protection truthtruths district turn smith review minute dutyduties survey presence influence stonestones dog benefit collection executive speechspeeches function queen marriage stock failure kitchen student effort holiday career attack length horse progressprogresses plant visit relation ball memorymemories bar opinion quarter impact scale race image trust justice edge gasgases railway expression advantage gold wood network text forest sister chair cause footfeet rise halfhalves winter corner insuranceinsurances step damage creditcredits pain possibilitypossibilities legislationlegislations strength speed crime hill debate will supplysupplies present confidence mary patient wind solution band museummuseums farm pound henryhenries matchmatches assessment message football nonoes animal skin scene article stuffstuffs introduction play administrationadministrations fear dad proportion island contact japan claim kingdom videovideos tvtvs existence telephone move traffic distance relief cabinet unemployment realityrealities target trial rock concept spirit accident organization construction coffee phone distribution train sight difficultydifficulties factor exercise weekend battle prison grant aircraftaircraft treetrees bridge strategystrategies contrast communication background shape wine star hope selection detail user path client searchsearches master rainrains offer goal dinner freedom attitude while agencyagencies seat manner favour fig.fig. pairpairs crisiscrises smile prince danger call capacitycapacities output note procedure theatre tour recognitionrecognitions middle absenceabsences sentence package track card sign commitment player threat weather element conflict notice victoryvictories bottom finance fund violence file profit standard jack route china expenditureexpenditures second discipline cell pp. reaction castle congresscongresses individual lead consideration debt option payment exhibition reform emphasisemphases spring audience feature touchtouches estate assemblyassemblies volume youth contribution curriculumcurricula appearance martin tom boat institute membership branchbranches busbuses waste heat neck object captain driver challenge conversation occasion code crown birth silencesilences literatureliteratures faith hell entryentries transfer gentlemangentlemen bag coal investigation leg belief total major document description murder aim manchester flight conclusion drug tradition pleasure connection owner treatytreaties tonytonies alan desiredesires professor copycopies ministryministries acid palace addressaddresses institution lunchlunches generation partner engine newspaper crosscrosses reduction welfarewelfares definition key release vote examination judge atmosphere leadership skyskies breath creation row guide milk cover screen intention criticism jonesjoneses silver customer journey explanation green measure brain significance phase injuryinjuries run coast technique valley drink magazine potential drive revolution bishop settlement christchrists metal motion indexindexes adult inflation sport surprise pension factoryfactories tape flow iron trip lane pool independence hole un flat content pay noise combination session appointment fashion consumer accommodation temperature mike religion author nation northern sample assistanceassistances interpretation aspect displaydisplays shoulder agent gallerygalleries republic cancer proposal sequence simon ship interview vehicle democracydemocracies improvement involvementinvolvements general enterprise van meal breakfast motor channel impression tone sheet pollution bob beautybeauties square vision spot distinction brown crowd fuel desk sum decline revenue fall diet bedroom soil reader shock fruit behalfbehalves deputydeputies roofroofs nose steel co artist graham plate song maintenancemaintenances formation grassgrasses spokesmanspokesmen ice talk program link ring expert establishment plastic candidate rail passage joe parishparishes ref emergencyemergencies liabilityliabilities identityidentities location framework strike countryside map lake household approval border bottle bird constitution autumn cat agriculture concentrationconcentrations guy dressdresses victim mountain editor theme error loan stressstresses recoveryrecoveries electricityelectricities recession wealthwealths request comparison lewislewises white walk focusfoci chief parent sleep massmasses jane bushbushes foundation bath item lifespan leelees publication decade beachbeaches sugar height charitycharities writer panel struggle dream outcome efficiencyefficiencies offenceoffences resolution reputation specialist taylortaylors pub co-operation port incident representation bread chain initiative clause resistance mistake worker advance empire notion mirror deliverydeliveries chest licencelicences frank average awarenessawarenesses travel expansion block alternative chancellor meat store selfselves break dramadramas corporation currencycurrencies extension convention partnership skill furniturefurnitures round regime inquiryinquiries rugbyrugbies philosophyphilosophies scope gate minorityminorities intelligence restaurant consequence mill golf retirement prioritypriorities plane gun gap core uncle thatcher fun arrival snow no.nos. command abuse limit championship"
)

var noun_vec [1000]string
var index = 0

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	questions := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		questions <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)

	go get_questions(questions, answers)
	go predictions(answers)
	go foretell_thy_future(answers)

	return questions
}

func get_questions(questions <-chan string, answers chan<- string) {
	for s := range questions {
		prophecy(s, answers)
	}

}

func predictions(predictions chan<- string) {
	nonsense := []string{
		"The sea",
		"The stars",
		"The ocean",
		"The mountains",
		"Mountains",
		"People",
		"",
	}
	for {
		time.Sleep(time.Duration(17+rand.Intn(20)) * time.Second)
		out := ""
		if index > 0 {
			out = nonsense[rand.Intn(len(nonsense))] + "..." + noun_vec[rand.Intn(index)]
		} else {
			out = nonsense[rand.Intn(len(nonsense))]
		}
		out += "?"
		prophecy(out, predictions)
	}
}

func foretell_thy_future(answers <-chan string) {
	for a := range answers {
		fmt.Printf("%s: ", star)
		for _, char := range a {
			fmt.Printf("%c", char)
			RandomSleep(15)
		}
		fmt.Printf("\n> ")
	}
}

func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Find the longest word.
	longestWord := ""
	noun := ""
	isThe := false
	if strings.Contains(question, "?") {
		question = strings.Replace(question, "?", "", 1)
		words := strings.Fields(question) // Fields extracts the words into a slice.
		for _, w := range words {
			if strings.Contains(strings.ToLower(w), "the") {
				noun = w + " "
				isThe = true
			} else if isThe == true {
				noun += w
				isThe = false
				if index > 999 {
					index = 0
				}
				noun_vec[index] = noun
				index += 1
				break
			} else if strings.Contains(nouns, strings.ToLower(w)) {
				noun = w
				if index > 999 {
					index = 0
				}
				noun_vec[index] = noun
				index += 1
			}
			if len(w) > len(longestWord) {
				longestWord = w
			}
		}

		// Cook up some pointless nonsense.
		nonsense := []string{
			"The moon is dark.",
			"The sun is bright.",
			"Good",
			"No good",
			"sus",
		}

		if len(noun) > 0 {
			answer <- noun + "... " + nonsense[rand.Intn(len(nonsense))]
		} else {
			answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
		}
	} else {
		answer <- "that is an interesting fact... Maybe?"
	}

}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}
