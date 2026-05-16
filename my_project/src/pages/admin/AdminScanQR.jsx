import React, { useEffect, useState } from "react";
import { Container, Card, Alert, Button } from "react-bootstrap";
import { Html5Qrcode } from "html5-qrcode";
import Cookies from "js-cookie";
import { useNavigate } from "react-router-dom";

export default function AdminScanQR() {

  const [message, setMessage] = useState(null);

  const token = Cookies.get("token");
  const navigate = useNavigate();

  const handleVisit = async (membershipId) => {

    try {

      const res = await fetch(
        `http://localhost:8080/api/v1/admin/visits/${membershipId}`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
          },
        }
      );

      const data = await res.json();

      if (!res.ok) {
        setMessage({ type: "danger", text: data.error });
        return;
      }

      setMessage({ type: "success", text: "บันทึกการเข้าใช้สำเร็จ" });

    } catch {
      setMessage({ type: "danger", text: "เกิดข้อผิดพลาดในการเชื่อมต่อ" });
    }

  };

  useEffect(() => {

    let scanner;
    let scanned = false; // กัน scan ซ้ำ

    const startScanner = async () => {

      try {

        scanner = new Html5Qrcode("qr-reader");

        await scanner.start(
          { facingMode: "environment" },
          {
            fps: 10,
            qrbox: 250
          },
          async (decodedText) => {

            if (scanned) return;
            scanned = true;

            console.log("SCAN:", decodedText);

            let membershipId = decodedText;

            if (decodedText.startsWith("FITNESS:")) {
              membershipId = decodedText.replace("FITNESS:", "");
            }

            await handleVisit(membershipId);

            // หยุดกล้องหลังสแกน
            if (scanner) {
              scanner.stop().catch(() => {});
            }

          }
        );

      } catch (err) {
        console.log("Camera error:", err);
      }

    };

    setTimeout(startScanner, 300);

    return () => {
      if (scanner) {
        scanner.stop().catch(() => {});
      }
    };
// eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return (

    <Container className="py-5">

      <Card className="shadow-lg p-4 text-center">

        <h3 className="mb-4">Scan QR Code</h3>

        <div id="qr-reader" style={{ width: "100%" }}></div>

        {message && (
          <Alert className="mt-3" variant={message.type}>
            {message.text}
          </Alert>
        )}

        <Button
          variant="secondary"
          className="mt-3"
          onClick={() => navigate("/admin/users")}
        >
          Back
        </Button>

      </Card>

    </Container>
  );
}