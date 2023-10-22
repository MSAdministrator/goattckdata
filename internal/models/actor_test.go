package models

import (
	"encoding/json"
	"testing"
)

var testActor = `
{
	"id": "intrusion-set--6a2e693f-24e5-451a-9f88-b36a108e5662",
	"name": "APT1",
	"created": "2017-05-31T21:31:47.955Z",
	"modified": "2021-05-26T12:23:48.842Z",
	"x_mitre_version": "1.4",
	"type": "intrusion-set",
	"aliases": [
		"APT1",
		"Comment Crew",
		"Comment Group",
		"Comment Panda"
	],
	"x_mitre_contributors": [],
	"revoked": false,
	"description": "[APT1](https://attack.mitre.org/groups/G0006) is a Chinese threat group that has been attributed to the 2nd Bureau of the People\u2019s Liberation Army (PLA) General Staff Department\u2019s (GSD) 3rd Department, commonly known by its Military Unit Cover Designator (MUCD) as Unit 61398. (Citation: Mandiant APT1)",
	"x_mitre_modified_by_ref": "identity--c78cb6e5-0c4b-4611-8297-d1b8b55e40b5",
	"x_mitre_deprecated": false,
	"x_mitre_attack_spec_version": "",
	"created_by_ref": "identity--c78cb6e5-0c4b-4611-8297-d1b8b55e40b5",
	"x_mitre_domains": [
		"enterprise-attack"
	],
	"object_marking_refs": [
		"marking-definition--fa42a846-8d90-4e51-bc29-71d5b4802168"
	],
	"external_references": [
		{
			"source_name": "mitre-attack",
			"url": "https://attack.mitre.org/groups/G0006",
			"external_id": "G0006",
			"description": ""
		},
		{
			"source_name": "APT1",
			"url": "",
			"external_id": "",
			"description": "(Citation: Mandiant APT1)"
		},
		{
			"source_name": "Comment Crew",
			"url": "",
			"external_id": "",
			"description": "(Citation: Mandiant APT1)"
		},
		{
			"source_name": "Comment Group",
			"url": "",
			"external_id": "",
			"description": "(Citation: Mandiant APT1)"
		},
		{
			"source_name": "Comment Panda",
			"url": "",
			"external_id": "",
			"description": "(Citation: CrowdStrike Putter Panda)"
		},
		{
			"source_name": "Mandiant APT1",
			"url": "https://www.fireeye.com/content/dam/fireeye-www/services/pdfs/mandiant-apt1-report.pdf",
			"external_id": "",
			"description": "Mandiant. (n.d.). APT1 Exposing One of China\u2019s Cyber Espionage Units. Retrieved July 18, 2016."
		},
		{
			"source_name": "CrowdStrike Putter Panda",
			"url": "http://cdn0.vox-cdn.com/assets/4589853/crowdstrike-intelligence-report-putter-panda.original.pdf",
			"external_id": "",
			"description": "Crowdstrike Global Intelligence Team. (2014, June 9). CrowdStrike Intelligence Report: Putter Panda. Retrieved January 22, 2016."
		}
	],
	"names": [
		"Comment Crew",
		"Comment Panda",
		"PLA Unit 61398",
		"TG-8223",
		"APT1",
		"BrownFox",
		"GIF89a, ShadyRAT, Shanghai Group, Byzantine Candor",
		"G0006"
	],
	"external_tools": [
		"WEBC2",
		"BISCUIT and many others"
	],
	"country": [],
	"operations": [],
	"links": [
		"http://www.mcafee.com/us/resources/white-papers/wp-operation-shady-rat.pdf",
		"http://www.nytimes.com/2013/02/19/technology/chinas-army-is-seen-as-tied-to-hacking-against-us.html?emc=na&_r=2&",
		"https://www.secureworks.com/research/analysis-of-dhs-nccic-indicators",
		"https://www.scribd.com/doc/13731776/Tracking-GhostNet-Investigating-a-Cyber-Espionage-Network",
		"http://www.nartv.org/mirror/ghostnet.pdf"
	],
	"targets": [
		"U.S. cybersecurity firm Mandiant, later purchased by FireEye, released a report in February 2013 that exposed one of China's cyber espionage units, Unit 61398. The group, which FireEye called APT1, is a unit within China's People's Liberation Army (PLA) that has been linked to a wide range of cyber operations targeting U.S. private sector entities for espionage purposes. The comprehensive report detailed evidence connecting APT1 and the PLA, offered insight into APT1's operational malware and methodologies, and provided timelines of the espionage it conducted."
	],
	"external_description": [],
	"attck_id": "",
	"attck_ids": [],
	"comment": "",
	"comments": []
}`

func jsonToMap(jsonStr string) map[string]interface{} {
	result := make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &result)
	return result
}

func TestNewActor(t *testing.T) {
	actor, err := NewActor(jsonToMap(testActor))
	if err != nil {
		t.Errorf("Error, could not load Actor: %v", err)
	}
	if actor.Name != "APT1" {
		t.Errorf("Error, could not load Actor data models: %v", err)
	}
	if actor.XMitreVersion != "1.4" {
		t.Errorf("Error, could not load Actor data models: %v", err)
	}

}
