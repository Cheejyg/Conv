package com.github.cheejyg.conv.java;

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// import com.github.cheejyg.conv.java.*;
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/**
 * Long conversions.
 * 
 * @author Cheejyg
 */
public class Long {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Long conversions
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	/**
	 * LongToString parses <code>long</code> long_ to <code>String</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>String</code>
	 */
	public static java.lang.String LongToString(long long_) {
		java.lang.String string_ = Conv.ToString(long_);

		return string_;
	}

	/**
	 * LongToLong parses <code>long</code> long_ to <code>Long</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Long</code>
	 */
	public static java.lang.Long LongToLong(long long_) {
		return long_;
	}

	/**
	 * LongToInteger parses <code>long</code> long_ to <code>Integer</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Integer</code>
	 */
	public static java.lang.Integer LongToInteger(long long_) {
		java.lang.Integer integer_ = java.lang.Integer.valueOf((int) long_);

		return integer_;
	}

	/**
	 * LongToShort parses <code>long</code> long_ to <code>Short</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Short</code>
	 */
	public static java.lang.Short LongToShort(long long_) {
		java.lang.Short short_ = java.lang.Short.valueOf((short) long_);

		return short_;
	}

	/**
	 * LongToByte parses <code>long</code> long_ to <code>Byte</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Byte</code>
	 */
	public static java.lang.Byte LongToByte(long long_) {
		java.lang.Byte byte_ = java.lang.Byte.valueOf((byte) long_);

		return byte_;
	}

	/**
	 * LongToBoolean parses <code>long</code> long_ to <code>Boolean</code>.
	 * 
	 * @author Cheejyg
	 * @param long_ The <code>long</code> to be parsed.
	 * @return <code>Boolean</code>
	 */
	public static java.lang.Boolean LongToBoolean(long long_) {
		if (long_ == 0) {
			return false;
		} else if (long_ == 1) {
			return true;
		} else {
		}

		return null;
	}
}
