import QRCode from "react-qr-code";

export default function QRCodeMember({ membershipID }) {

  return (
    <div style={{
      background: "white",
      padding: "20px",
      textAlign: "center",
      width: "220px"
    }}>

      <h4>Member QR</h4>

      <QRCode
        value={membershipID.toString()}
        size={180}
      />

      <p>ID : {membershipID}</p>

    </div>
  );
}