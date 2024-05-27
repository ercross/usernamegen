package usernamegen

import (
	"fmt"
	"math/rand"
)

type Username struct {
	Username    string `bson:"username" sql:"username"`
	hasBeenUsed bool   `bson:"used" sql:"used"`
}

func generateNewBatch(count int64, separator string) []Username {
	batch := make([]Username, 0, count)
	for i := int64(0); i < count; i++ {
		batch[i] = generateNewUsername(separator)
	}
	return batch
}

func generateNewUsername(separator string) Username {
	leftIndex := rand.Intn(len(adjectives) - 1)
	rightIndex := rand.Intn(len(nouns) - 1)

	username := fmt.Sprintf("%s%s%s", adjectives[leftIndex], separator, nouns[rightIndex])
	return Username{
		Username:    username,
		hasBeenUsed: false,
	}
}

func totalPossibleCombinationsCount() int {
	return len(adjectives) * len(nouns)
}

var adjectives = []string{"adventurous", "affectionate", "ambitious", "artistic", "assertive", "bold", "brave", "bright", "calm", "charismatic", "clever", "compassionate", "confident", "curious", "daring", "dedicated", "delightful", "determined", "devoted", "diligent", "eager", "elegant", "empathetic", "energetic", "enthusiastic", "faithful", "fearless", "fierce", "forgiving", "friendly", "generous", "gentle", "graceful", "gracious", "grateful", "happy", "helpful", "honest", "humble", "humorous", "idealistic", "imaginative", "inspirational", "intelligent", "intuitive", "inventive", "joyful", "jolly", "jubilant", "keen", "kind", "kindhearted", "knowledgeable", "lighthearted", "lively", "lovable", "loving", "loyal", "majestic", "mighty", "mindful", "motivated", "modest", "noble", "nurturing", "observant", "open-minded", "optimistic", "outgoing", "passionate", "patient", "peaceful", "persistent", "quick", "quick-witted", "quiet", "quirky", "radiant", "reliable", "resilient", "resourceful", "respectful", "sincere", "spirited", "strong", "supportive", "sympathetic", "talented", "tenacious", "thoughtful", "tolerant", "trustworthy", "understanding", "unique", "upbeat", "valiant", "versatile", "vibrant", "virtuous", "warm", "warm-hearted", "welcoming", "wise", "witty", "xenial", "xenodochial", "yearning", "youthful", "zany", "zealous", "zestful", "zesty"}
var nouns = []string{"apple", "adventure", "armor", "anchor", "archipelago", "atom", "acorn", "arch", "bridge", "breeze", "bay", "book", "cloud", "canyon", "castle", "desert", "dragon", "echo", "earth", "forest", "fountain", "garden", "galaxy", "grove", "glacier", "glimmer", "harbor", "hill", "island", "jungle", "knoll", "kite", "lighthouse", "lake", "meadow", "mountain", "night", "ocean", "palace", "pyramid", "quasar", "rainbow", "river", "reef", "rock", "shore", "stream", "star", "sunset", "sky", "temple", "tower", "valley", "vista", "wind", "yard", "zenith", "zenith", "zenith", "zephyr", "xylophone", "yarn", "xylophone", "yard", "yard", "zenith", "anchor", "beach", "bridge", "cave", "delta", "dune", "elv", "fir", "fjord", "horizon", "kelp", "lava", "moss", "marsh", "nest", "oasis", "peak", "quarry", "reef", "shore", "stream", "thicket", "undergrowth", "vine", "water", "xenolith", "yard", "zenith", "archipelago", "bay", "cave", "delta", "dune", "earth", "forest", "grove", "hill", "ice", "jungle", "kelp", "lava", "moss", "marsh", "nest", "oasis", "peak", "quarry", "reef", "stream", "thicket", "undergrowth", "vine", "waterfall", "xylem", "yard", "zenith", "atom", "breeze", "canyon", "dust", "echo", "flame", "glimmer", "haze", "insect", "light", "mist", "particle", "ray", "spectrum", "tide", "universe", "vapor", "wave", "yard", "yarn", "zenith", "arch", "bridge", "chamber", "dome", "elevation", "fountain", "garden", "hall", "igloo", "jewel", "knob", "lantern", "mansion", "obelisk", "pavilion", "quoin", "roof", "spire", "tower", "urn", "vault", "window", "xyst", "yard", "ziggurat", "zeus", "hera", "poseidon", "hades", "apollo", "artemis", "ares", "aphrodite", "hermes", "athena", "dionysus", "demeter", "hephaestus", "hestia", "persephone", "pan", "selene", "eos", "nyx", "chaos", "uranus", "gaea", "cronus", "rhea", "prometheus", "atlas", "heracles", "theseus", "perseus", "odysseus", "achilles", "hector", "orpheus", "eurydice", "icarus", "daedalus", "sisyphus", "midas", "medusa", "chiron", "cerberus", "pegasus", "sphinx", "minotaur", "cyclops", "sirens", "centaur", "satyr", "nymph", "dryad", "naiad", "furies", "harpy", "chimera", "hydra", "griffin", "phoenix", "hurricane", "earthquake", "winter", "autumn", "spring", "summer", "tornado", "flood", "drought", "blizzard", "hailstorm", "tsunami", "avalanche", "volcano", "wildfire", "cyclone", "thunderstorm", "lightning", "rain", "snow", "fog", "dew", "frost", "sleet", "hail", "eclipse", "solstice", "equinox", "aurora", "comet", "meteor", "constellation", "zodiac", "capricorn", "aquarius", "pisces", "aries", "taurus", "gemini", "cancer", "leo", "virgo", "libra", "scorpio", "sagittarius", "mercury", "venus", "earth", "mars", "jupiter", "saturn", "uranus", "neptune", "pluto", "superman", "batman", "thor", "hulk", "vision", "ant-man", "wasp", "aquaman", "flash", "cyborg", "shazam", "nightwing", "starfire", "raven", "rogue", "gambit", "iceman", "colossus", "nova", "speed", "patriot", "stature", "sentry", "loki", "heimdall", "korg", "rocket", "groot", "mantis", "drax", "nebula", "yondu", "shuri", "okoye", "nakia", "abomination", "dormammu", "venom", "sandman", "mysterio", "vulture", "electro", "lizard", "kingpin", "bullseye", "mojo", "ant", "bat", "bear", "bee", "bird", "cat", "cow", "crab", "deer", "dog", "dolphin", "duck", "eagle", "fish", "fox", "frog", "goat", "goose", "hawk", "horse", "lion", "lizard", "moose", "mouse", "owl", "panda", "pig", "rabbit", "rat", "seal", "shark", "sheep", "snake", "swan", "tiger", "toad", "turtle", "whale", "wolf", "zebra"}
