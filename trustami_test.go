package ai_connector_go

import "testing"

func TestLangDetector_DetectLanguage(t *testing.T) {
	langDetector := NewLangDetector(TestToken)

	lang, err := langDetector.DetectLanguage("This is a test")
	if err != nil {
		t.Error(err)
	}

	if lang.Language != "en" {
		t.Error("Language should be 'en'")
	}
}

func TestSentimentAnalyser_AnalyseSentiment(t *testing.T) {
	sentimentAnalyser := NewSentimentAnalyser(TestToken)

	sentiment, err := sentimentAnalyser.AnalyseSentiment("I hate this product")
	if err != nil {
		t.Error(err)
	}

	if sentiment.Polarity != "negative" {
		t.Error("Polarity should be 'negative'")
	}
}

func TestSummarizer_SummarizeText(t *testing.T) {
	summarizer := NewSummarizer(TestToken)

	summary, err := summarizer.SummarizeText("Nach dem Einmarsch russischer Truppen in die Ostukraine hat Ukraines Präsident Wolodymyr Selenskyj eine allgemeine Mobilmachung angeordnet. Das Staatsoberhaupt habe ein entsprechendes Dekret unterschrieben, meldete die Agentur Unian unter Berufung auf das Präsidialamt in Kiew. Die Anordnung gilt demnach 90 Tage und sieht die Einberufung von Wehrpflichtigen und Reservisten vor.", 1)
	if err != nil {
		t.Error(err)
	}

	if summary.Summary == "" {
		t.Error("Summary should not be empty")
	}
}
