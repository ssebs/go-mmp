// arduino-mmp.ino
// Source code for the arduino that will be communicating over Serial to the Go-MMP driver.
// If you're copying this yourself, you may need to tweak the BTN_PINs to match your pins.

#define LED_PIN 17
#define BTN_1_PIN 5
#define BTN_2_PIN 7
#define BTN_3_PIN 9
#define BTN_4_PIN 8
#define _DELAY 150

void setup() {
  // pinMode(LED_PIN, OUTPUT);
  pinMode(BTN_1_PIN, INPUT_PULLUP);
  pinMode(BTN_2_PIN, INPUT_PULLUP);
  pinMode(BTN_3_PIN, INPUT_PULLUP);
  pinMode(BTN_4_PIN, INPUT_PULLUP);
  
  Serial.begin(9600);
}

void loop() {
  handleInput();
}

void handleInput() {
  
  if (digitalRead(BTN_1_PIN) == LOW) {
    Serial.print("0");
    delay(_DELAY);
  }
  if (digitalRead(BTN_2_PIN) == LOW) {
    Serial.print("1");
    delay(_DELAY);
  }
  if (digitalRead(BTN_3_PIN) == LOW) {
    Serial.print("2");
    delay(_DELAY);
  }
  if (digitalRead(BTN_4_PIN) == LOW) {
    Serial.print("3");
    delay(_DELAY);
  }
}

void blink() {
  digitalWrite(LED_PIN, HIGH);
  delay(500);
  digitalWrite(LED_PIN, LOW);
  delay(500);
}